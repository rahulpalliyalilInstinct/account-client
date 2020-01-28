package behaviour

import (
	"fmt"
	"time"

	"account-client/pkg/deregistering"
	"account-client/pkg/http"
	"account-client/pkg/listing"
	"account-client/pkg/registering"
)

type pageNum int

type TestListAccountClient struct {
	client              *http.AccountClient
	identifier          *testIdentifier
	accountIDs          []string
	expectedAccountsMap map[pageNum]*listing.Accounts
}

func NewListTestClient() *TestListAccountClient {
	client, err := http.NewClient(int(5*time.Second), http.Config{})
	if err != nil {
		return nil
	}
	return &TestListAccountClient{
		client:     client,
		identifier: &testIdentifier{},
	}
}

func (t *TestListAccountClient) multipleAccountsWithAndWithinSameCountry(accountID1, accountID2, countryCode string) error {
	t.accountIDs = []string{accountID1, accountID2}
	for _, accountID := range t.accountIDs {
		account := registering.Account{
			Data: registering.Data{
				ID: accountID,
				Attributes: registering.Attributes{
					Country: countryCode,
				},
			},
		}
		t.identifier.id = accountID
		service := registering.NewService(t.client, t.identifier)
		if err := service.CreateAccount(account); err != nil {
			return err
		}
	}
	return nil
}

func (t *TestListAccountClient) iWantToListASingleAccountPerPage() error {
	t.expectedAccountsMap = make(map[pageNum]*listing.Accounts)
	service := listing.NewService(t.client)
	for i := 0; i < 2; i++ {
		accounts, err := service.GetAccounts(listing.Page{
			Number: i,
			Size:   1,
		})
		if err != nil {
			return fmt.Errorf("getAccounts() request failed: %v", err)
		}
		t.expectedAccountsMap[pageNum(i)] = accounts
	}
	return nil
}

func (t *TestListAccountClient) iAmAbleToSeeMyFirstAccountInPageAndSecondAccountInPage() error {
	if len(t.expectedAccountsMap) != 2 {
		return fmt.Errorf("listAccounts(): mismatch in accounts found: %v", t.expectedAccountsMap)
	}
	for pageNum, accounts := range t.expectedAccountsMap {
		if len(accounts.Data) != 1 {
			continue
		}

		if t.accountIDs[int(pageNum)] != accounts.Data[0].ID {
			return fmt.Errorf("expected account id: '%v', got: '%v'", t.accountIDs[int(pageNum)], accounts.Data[0].ID)
		}
	}
	return nil
}

func (t *TestListAccountClient) cleanUp() {
	service := deregistering.NewService(t.client)
	for _, accountID := range t.accountIDs {
		service.DeleteAccount(deregistering.Account{
			ID: accountID,
		})
	}
}
