package main

import (
	"time"

	"account-client/pkg/deregistering"
	"account-client/pkg/http"
	"account-client/pkg/listing"
	"account-client/pkg/registering"
	"account-client/pkg/uuid"
)

func main() {
	var id string
	client, err := http.NewClient(int(5*time.Second), http.Config{})
	if err != nil {
		return
	}
	// example for registering/creating an account
	for i := 0; i < 2; i++ {
		uniqueIdentifier := uuid.NewUniqueIdentifier()
		service := registering.NewService(client, uniqueIdentifier)
		if err = service.CreateAccount(registering.Account{Data: registering.Data{Attributes: registering.Attributes{Country: "GB"}}}); err != nil {
			return
		}
	}

	// example for listing accounts
	listingService := listing.NewService(client)
	count := 0
	for {
		accounts, listingErr := listingService.GetAccounts(*listing.NewPage(count, 1))
		if listingErr != nil {
			return
		}
		if accounts.Links.Next == "" {
			id = accounts.Data[0].ID
			break
		}
		count++
	}

	// example for fetching an account
	account, err := listingService.GetAccount(id)
	if err != nil {
		return
	}
	// example for deregistering/removing an account
	deRegService := deregistering.NewService(client)
	err = deRegService.DeleteAccount(deregistering.Account{
		ID: account.Data.ID,
	})
	if err != nil {
		return
	}
}
