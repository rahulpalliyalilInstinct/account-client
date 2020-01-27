package behaviour_tests

import (
	"fmt"
	"time"

	"account-client/pkg/deregistering"
	"account-client/pkg/http"
	"account-client/pkg/listing"
	"account-client/pkg/registering"
)

type testRegisterAccountClient struct {
	account    registering.Account
	client     *http.AccountClient
	identifier *testIdentifier
}

type testIdentifier struct {
	id string
}

func (t *testIdentifier) Generator() string {
	return t.id
}

func NewRegisterTestClient() *testRegisterAccountClient {
	client, err := http.NewClient(int(5*time.Second), http.Config{})
	if err != nil {
		return nil
	}
	return &testRegisterAccountClient{
		client:     client,
		identifier: &testIdentifier{},
	}
}

func (t *testRegisterAccountClient) aNewUnregisteredAccountWithCountryCode(id, countryCode string) error {
	t.account = registering.Account{
		Data: registering.Data{
			ID: id,
			Attributes: registering.Attributes{
				Country: countryCode,
			},
		},
	}
	t.identifier.id = id
	return nil
}

func (t *testRegisterAccountClient) iSendARequestToRegisterTheAccount() error {
	service := registering.NewService(t.client, t.identifier)
	if err := service.CreateAccount(t.account); err != nil {
		return err
	}
	return nil
}

func (t *testRegisterAccountClient) iAmAbleToSeeMyAccountRegistered() error {
	service := listing.NewService(t.client)
	account, err := service.GetAccount(t.account.ID)
	if err != nil {
		return fmt.Errorf("GetAccount() request failed: %v", err)
	}

	if account.Data.ID != t.account.ID {
		return fmt.Errorf("expected account id: '%v', got: '%v'", t.account.ID, account.Data.ID)
	}

	return nil

}

func (t *testRegisterAccountClient) cleanUp() {
	service := deregistering.NewService(t.client)
	service.DeleteAccount(deregistering.Account{
		ID: t.account.ID,
	})
}
