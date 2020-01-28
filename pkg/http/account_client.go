package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"account-client/pkg/deregistering"
	"account-client/pkg/listing"
	"account-client/pkg/registering"
	"github.com/pkg/errors"
)

const (
	createAccountReqFailedErrMsg = "create account request failed"
	deleteAccountReqFailedErrMsg = "delete account request failed"
	fetchAccountReqFailedErrMsg  = "fetch account request failed"
	listAccountsReqFailedErrMsg  = "list accounts request failed"

	accountsApi    = "/v1/organisation/accounts/"
	defaultAddress = "http://localhost:8080"
)

type QueryParams map[string]string
type Header map[string]string

type Error struct {
	StatusCode int
	Message    string
}

func (e *Error) Error() string {
	return fmt.Sprintf("HTTPStatus: %v, Message: %q",
		e.StatusCode, e.Message)
}

// Config defines configuration parameters for a new client.
type Config struct {
	// The address of the accountapi to connect to.
	Address string
}

type AccountClient struct {
	timeout    time.Duration
	httpClient *http.Client
	endpoint   *url.URL
}

// NewClient provides a new accountClient.If there is no address
// of the accountapi specified, then the default address is used.
func NewClient(timeoutInSeconds int, cfg Config) (*AccountClient, error) {
	if cfg.Address == "" {
		cfg.Address = defaultAddress
	}
	u, err := url.Parse(cfg.Address)
	if err != nil {
		return nil, err
	}
	u.Path = strings.TrimRight(u.Path, "/")
	u.Path += accountsApi

	return &AccountClient{
		endpoint:   u,
		httpClient: &http.Client{},
		timeout:    time.Second * time.Duration(timeoutInSeconds),
	}, nil
}

// Create creates an account. if creation fails then a non nil error
// is returned.
func (c *AccountClient) Create(account registering.Account) error {
	bodyBytes, err := json.Marshal(account)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.endpoint.String(), bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}

	respBody, err := c.request(context.Background(), req, nil, nil)
	if err != nil {
		return errors.Wrap(err, createAccountReqFailedErrMsg)
	}
	defer respBody.Close()
	return nil
}

// Delete deletes an account. if deletion fails then a non nil error
// is returned.
func (c *AccountClient) Delete(account deregistering.Account) error {
	endpointUrl := c.endpoint.String() + account.ID
	req, err := http.NewRequest("DELETE", endpointUrl, nil)
	if err != nil {
		return nil
	}

	queryParams := make(map[string]string)
	queryParams["version"] = account.Version
	respBody, err := c.request(context.Background(), req, nil, queryParams)
	if err != nil {
		return errors.Wrap(err, deleteAccountReqFailedErrMsg)
	}
	defer respBody.Close()
	return nil
}

// Fetch gets an account , given the accountId. If an invalid
// id is provided, then a non nil error is returned.
func (c *AccountClient) Fetch(id string) (*listing.Account, error) {
	account := &listing.Account{}
	endpointUrl := c.endpoint.String() + id
	req, err := http.NewRequest("GET", endpointUrl, nil)
	if err != nil {
		return nil, err
	}

	respBody, err := c.request(context.Background(), req, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, fetchAccountReqFailedErrMsg)
	}
	defer respBody.Close()
	responseDecoder := json.NewDecoder(respBody)
	if err := responseDecoder.Decode(account); err != nil {
		return nil, err
	}
	return account, nil
}

// List lists the accounts based on the page number and page size.
// An error is returned , if there is a failure in retrieving the accounts.
func (c *AccountClient) List(page listing.Page) (*listing.Accounts, error) {
	accounts := &listing.Accounts{}
	req, err := http.NewRequest("GET", c.endpoint.String(), nil)
	if err != nil {
		return nil, err
	}

	queryParams := make(map[string]string)
	queryParams["page[number]"] = strconv.Itoa(page.Number)
	queryParams["page[size]"] = strconv.Itoa(page.Size)

	respBody, err := c.request(context.Background(), req, nil, queryParams)
	if err != nil {
		return nil, errors.Wrap(err, listAccountsReqFailedErrMsg)
	}
	defer respBody.Close()
	responseDecoder := json.NewDecoder(respBody)
	if err := responseDecoder.Decode(accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}

func (c *AccountClient) request(ctx context.Context, req *http.Request, headers Header, params QueryParams) (io.ReadCloser, error) {
	for header, value := range headers {
		req.Header.Add(header, value)
	}

	q := req.URL.Query()
	for param, value := range params {
		q.Add(param, value) // Add a new value to the set.
	}
	req.URL.RawQuery = q.Encode()
	_, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, &Error{StatusCode: resp.StatusCode, Message: fmt.Sprintf("unexpected status code, %d", resp.StatusCode)}
	}
	return resp.Body, nil
}
