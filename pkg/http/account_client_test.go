package http_test

import (
	"strings"
	"testing"
	"time"

	"account-client/pkg/deregistering"
	"account-client/pkg/http"
	"account-client/pkg/listing"
	"account-client/pkg/registering"

	"github.com/stretchr/testify/assert"
	"gopkg.in/h2non/gock.v1"
)

func interceptGetAccountsCall(status int, expectedAccount string) {
	gock.New("").
		Get("/v1/organisation/accounts").
		Reply(status).
		JSON(expectedAccount)
}

func interceptFetchAccountCall(status int, expectedAccount string) {
	gock.New("").
		Get("/v1/organisation/accounts/*").
		Reply(status).
		JSON(expectedAccount)
}

func interceptCreateAccountCall(status int) {
	gock.New("").
		Post("/v1/organisation/accounts/").
		Reply(status)
}

func interceptDeleteAccountCall(status int) {
	gock.New("").
		Delete("/v1/organisation/accounts/*").
		Reply(status)
}

func TestAccountClient_Create(t *testing.T) {
	type fields struct {
	}
	type args struct {
		account registering.Account
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		fn         func(int)
		statusCode int
		wantErr    bool
		errMsg     string
	}{
		{
			name: "failed to create account",
			args: args{account: registering.Account{
				Data: registering.Data{
					Type:           "some type",
					ID:             "some id",
					OrganisationID: "organisation id",
					Attributes:     registering.Attributes{},
				},
			}},
			fn:         interceptCreateAccountCall,
			statusCode: 400,
			wantErr:    true,
			errMsg:     "create account request failed",
		},
		{
			name: "successfully create account",
			args: args{account: registering.Account{
				Data: registering.Data{
					Type:           "some type",
					ID:             "some id",
					OrganisationID: "organisation id",
					Attributes:     registering.Attributes{},
				},
			}},
			fn:         interceptCreateAccountCall,
			statusCode: 200,
			wantErr:    false,
			errMsg:     "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			test.fn(test.statusCode)
			defer gock.Off()
			client, _ := http.NewClient(int(5*time.Second), http.Config{})
			err := client.Create(registering.Account{})

			if test.wantErr && !strings.Contains(err.Error(), test.errMsg) {
				t.Errorf("Create() %s: got = %v, want = %v", test.name, err, test.errMsg)
				return
			}
			if !test.wantErr {
				assert.Nil(t, err)
			}
		})
	}
}

func TestAccountClient_Delete(t *testing.T) {
	type fields struct {
	}
	type args struct {
		account deregistering.Account
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		fn         func(int)
		statusCode int
		wantErr    bool
		errMsg     string
	}{
		{
			name:       "failed to delete account",
			args:       args{},
			fn:         interceptDeleteAccountCall,
			statusCode: 400,
			wantErr:    true,
			errMsg:     "delete account request failed",
		},
		{
			name:       "successfully deleted account",
			args:       args{account: deregistering.Account{}},
			fn:         interceptDeleteAccountCall,
			statusCode: 200,
			wantErr:    false,
			errMsg:     "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			test.fn(test.statusCode)
			defer gock.Off()
			client, _ := http.NewClient(int(5*time.Second), http.Config{})
			err := client.Delete(test.args.account)

			if test.wantErr && !strings.Contains(err.Error(), test.errMsg) {
				t.Errorf("Delete() %s: got = %v, want = %v", test.name, err, test.errMsg)
				return
			}
			if !test.wantErr {
				assert.Nil(t, err)
			}
		})
	}
}

func TestAccountClient_Fetch(t *testing.T) {
	const expectedID = "some id"
	expectedAccount := `{
  "data": {
    "id": "some id",
    "attributes": {
      "alternative_bank_account_names": null,
      "country": "GB"
    }
  }
}`
	type fields struct {
		expectedAccount string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		fn         func(int, string)
		statusCode int
		wantErr    bool
		errMsg     string
	}{
		{
			name:       "failed to fetch account",
			args:       args{id: "some id"},
			fields:     fields{expectedAccount: expectedAccount},
			fn:         interceptFetchAccountCall,
			statusCode: 400,
			wantErr:    true,
			errMsg:     "fetch account request failed",
		},
		{
			name:       "successfully fetched account",
			args:       args{id: "some id"},
			fields:     fields{expectedAccount: expectedAccount},
			fn:         interceptFetchAccountCall,
			statusCode: 200,
			wantErr:    false,
			errMsg:     "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			test.fn(test.statusCode, test.fields.expectedAccount)
			defer gock.Off()
			client, _ := http.NewClient(int(5*time.Second), http.Config{})
			account, err := client.Fetch(test.args.id)

			if test.wantErr && !strings.Contains(err.Error(), test.errMsg) {
				t.Errorf("Fetch() %s: got = %v, want = %v", test.name, err, test.errMsg)
				return
			}
			if !test.wantErr {
				assert.NotNil(t, account)
				assert.Equal(t, expectedID, account.Data.ID)
			}
		})
	}
}

func TestAccountClient_List(t *testing.T) {
	expectedID := "some id"
	expectedAccount := `{
  "data": [{
    "id": "some id",
    "attributes": {
      "alternative_bank_account_names": null,
      "country": "GB"
    }
  }]
}`
	type fields struct {
		expectedAccount string
	}
	type args struct {
		page listing.Page
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		fn         func(int, string)
		statusCode int
		wantErr    bool
		errMsg     string
	}{
		{
			name:       "failed to list accounts",
			args:       args{page: listing.Page{Number: 1, Size: 10}},
			fields:     fields{expectedAccount: expectedAccount},
			fn:         interceptGetAccountsCall,
			statusCode: 400,
			wantErr:    true,
			errMsg:     "list accounts request failed",
		},

		{
			name:       "successfully listed accounts",
			args:       args{page: listing.Page{Number: 1, Size: 10}},
			fields:     fields{expectedAccount: expectedAccount},
			fn:         interceptGetAccountsCall,
			statusCode: 200,
			wantErr:    false,
			errMsg:     "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			test.fn(test.statusCode, test.fields.expectedAccount)
			defer gock.Off()
			client, _ := http.NewClient(int(5*time.Second), http.Config{})
			accounts, err := client.List(test.args.page)

			if test.wantErr && !strings.Contains(err.Error(), test.errMsg) {
				t.Errorf("Fetch() %s: got = %v, want = %v", test.name, err, test.errMsg)
				return
			}
			if !test.wantErr {
				assert.NotNil(t, accounts)
				assert.Equal(t, expectedID, accounts.Data[0].ID)
			}
		})
	}
}
