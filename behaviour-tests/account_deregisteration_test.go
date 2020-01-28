package behaviour

import (
	"fmt"
	"time"

	"account-client/pkg/deregistering"
	"account-client/pkg/http"
	"account-client/pkg/listing"
	"account-client/pkg/registering"
)

type TestDeRegisterAccountClient struct {
	account    deregistering.Account
	client     *http.AccountClient
	identifier *testIdentifier
}

func NewTestDeRegisterClient() *TestDeRegisterAccountClient {
	client, err := http.NewClient(int(5*time.Second), http.Config{})
	if err != nil {
		return nil
	}
	return &TestDeRegisterAccountClient{
		client:     client,
		identifier: &testIdentifier{},
	}
}

func (t *TestDeRegisterAccountClient) aNewRegisteredAccountWithCountryCode(id, countryCode string) error {
	account := registering.Account{
		Data: registering.Data{
			ID: id,
			Attributes: registering.Attributes{
				Country: countryCode,
			},
		},
	}
	t.identifier.id = id
	service := registering.NewService(t.client, t.identifier)
	if err := service.CreateAccount(account); err != nil {
		return err
	}
	return nil
}

func (t *TestDeRegisterAccountClient) iSendARequestToDeregisterTheAccount() error {
	t.account = deregistering.Account{
		ID: t.identifier.id,
	}
	service := deregistering.NewService(t.client)
	return service.DeleteAccount(t.account)
}

func (t *TestDeRegisterAccountClient) iAmAbleToSeeMyAccountDeregistered() error {
	service := listing.NewService(t.client)
	account, err := service.GetAccount(t.account.ID)
	if account == nil {
		return nil
	}
	return fmt.Errorf("account is not deregistered: %v", err)
}
