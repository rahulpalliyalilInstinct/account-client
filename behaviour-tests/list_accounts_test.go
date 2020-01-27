package behaviour_tests

import (
	"fmt"
	"time"

	"account-client/pkg/deregistering"
	"account-client/pkg/http"
	"account-client/pkg/listing"
	"account-client/pkg/registering"
)

type pageNum int

type testListAccountClient struct {
	account             registering.Account
	client              *http.AccountClient
	identifier          *testIdentifier
	accountIds          []string
	expectedAccountsMap map[pageNum]*listing.Accounts
}

func NewListTestClient() *testListAccountClient {
	client, err := http.NewClient(int(5*time.Second), http.Config{})
	if err != nil {
		return nil
	}
	return &testListAccountClient{
		client:     client,
		identifier: &testIdentifier{},
	}
}

func (t *testListAccountClient) multipleAccountsWithAndWithinSameCountry(accountId1, accountid2, countryCode string) error {
	t.accountIds = []string{accountId1, accountid2}
	for _, accountId := range t.accountIds {
		account := registering.Account{
			Data: registering.Data{
				ID: accountId,
				Attributes: registering.Attributes{
					Country: countryCode,
				},
			},
		}
		t.identifier.id = accountId
		service := registering.NewService(t.client, t.identifier)
		if err := service.CreateAccount(account); err != nil {
			return err
		}
	}
	return nil
}

func (t *testListAccountClient) iWantToListASingleAccountPerPage() error {
	t.expectedAccountsMap = make(map[pageNum]*listing.Accounts)
	service := listing.NewService(t.client)
	for i := 0; i < 2; i++ {
		accounts, err := service.GetAccounts(listing.Page{
			Number: i,
			Size:   1,
		})
		if err != nil {
			return fmt.Errorf("GetAccounts() request failed: %v", err)
		}
		t.expectedAccountsMap[pageNum(i)] = accounts
	}
	return nil
}

func (t *testListAccountClient) iAmAbleToSeeMyFirstAccountInPageAndSecondAccountInPage() error {
	if len(t.expectedAccountsMap) != 2 {
		fmt.Errorf("ListAccounts(): mismatch in accounts found: %v", t.expectedAccountsMap)
	}
	for pageNum, accounts := range t.expectedAccountsMap {
		if len(accounts.Data) != 1 {
			continue
		}

		if t.accountIds[int(pageNum)] != accounts.Data[0].ID {
			return fmt.Errorf("expected account id: '%v', got: '%v'", t.accountIds[int(pageNum)], accounts.Data[0].ID)
		}
	}
	return nil
}

func (t *testListAccountClient) cleanUp() {
	service := deregistering.NewService(t.client)
	for _, accountId := range t.accountIds {
		service.DeleteAccount(deregistering.Account{
			ID: accountId,
		})
	}
}
