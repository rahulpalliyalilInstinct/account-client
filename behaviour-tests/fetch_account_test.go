package behaviour

import (
	"fmt"
	"time"

	"account-client/pkg/deregistering"
	"account-client/pkg/http"
	"account-client/pkg/listing"
)

type TestFetchAccountClient struct {
	account listing.Account
	client  *http.AccountClient
}

func NewFetchTestClient() *TestFetchAccountClient {
	client, err := http.NewClient(int(5*time.Second), http.Config{})
	if err != nil {
		return nil
	}
	return &TestFetchAccountClient{
		client: client,
	}
}

func (t *TestFetchAccountClient) iAmAbleToSeeMyAccountDetails(id string) error {
	service := listing.NewService(t.client)
	account, err := service.GetAccount(id)
	if err != nil {
		return fmt.Errorf("getAccount() request failed: %v", err)
	}
	if account == nil {
		return fmt.Errorf("getAccount() no record found for account id: %v", account.Data.ID)
	}
	t.account.Data.ID = account.Data.ID
	return nil
}

func (t *TestFetchAccountClient) cleanUp() {
	service := deregistering.NewService(t.client)
	service.DeleteAccount(deregistering.Account{
		ID: t.account.Data.ID,
	})
}
