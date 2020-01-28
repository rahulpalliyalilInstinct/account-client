package listing_test

import (
	"errors"
	"strings"
	"testing"

	"account-client/pkg/listing"

	"github.com/stretchr/testify/assert"
)

type testAccountClient struct {
	err      error
	account  *listing.Account
	accounts *listing.Accounts
}

func (t *testAccountClient) Fetch(id string) (*listing.Account, error) {
	return t.account, t.err
}

func (t *testAccountClient) List(page listing.Page) (*listing.Accounts, error) {
	return t.accounts, t.err
}

func TestNewService(t *testing.T) {
	t.Parallel()
	type fields struct {
		client listing.AccountAPIClient
	}
	tests := []struct {
		name          string
		fields        fields
		invalidHandle bool
	}{
		{
			name:          "returns nil service handle when a nil account client is passed",
			fields:        fields{},
			invalidHandle: true,
		},
		{
			name: "returns valid service",
			fields: fields{
				client: &testAccountClient{},
			},
			invalidHandle: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newService := listing.NewService(test.fields.client)

			if test.invalidHandle {
				assert.Nil(t, newService)
				return
			}
			assert.NotNil(t, newService)
		})
	}
}

func TestService_GetAccount(t *testing.T) {
	t.Parallel()
	const expectedErrMsg = "something bad happened"
	expectedAccount := listing.Account{
		Data: listing.Data{
			ID:   "some id",
			Type: "some Type",
		}}
	type fields struct {
		client listing.AccountAPIClient
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "invalid account id specified",
			fields:  fields{client: &testAccountClient{}},
			args:    args{id: ""},
			wantErr: true,
			errMsg:  "failed to specify valid account id",
		},
		{
			name:    "failed to fetch account",
			fields:  fields{client: &testAccountClient{err: errors.New(expectedErrMsg)}},
			args:    args{id: "valid id"},
			wantErr: true,
			errMsg:  expectedErrMsg,
		},
		{
			name:    "account retrieved successfully",
			fields:  fields{client: &testAccountClient{account: &expectedAccount}},
			args:    args{id: "valid id"},
			wantErr: false,
			errMsg:  "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newService := listing.NewService(test.fields.client)
			account, err := newService.GetAccount(test.args.id)

			if test.wantErr && !strings.Contains(err.Error(), test.errMsg) {
				t.Errorf("GetAccount() %s: got = %v, want = %v", test.name, err, test.errMsg)
				return
			}
			if !test.wantErr {
				assert.Equal(t, expectedAccount.Data.ID, account.Data.ID)
				assert.Equal(t, expectedAccount.Data.Type, account.Data.Type)
			}
		})
	}
}

func TestService_GetAccounts(t *testing.T) {
	t.Parallel()
	expectedErrMsg := "something bad happened"
	expectedAccounts := listing.Accounts{
		Data: []listing.Data{{
			ID:   "some id",
			Type: "some Type",
		}}}
	type fields struct {
		client listing.AccountAPIClient
	}
	type args struct {
		page listing.Page
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "invalid page size is specified",
			fields:  fields{client: &testAccountClient{}},
			args:    args{page: listing.Page{Size: 0}},
			wantErr: true,
			errMsg:  "failed to provide valid page size",
		},
		{
			name:    "failed to list accounts",
			fields:  fields{client: &testAccountClient{err: errors.New(expectedErrMsg)}},
			args:    args{page: listing.Page{Size: 1}},
			wantErr: true,
			errMsg:  expectedErrMsg,
		},
		{
			name:    "successfully list accounts",
			fields:  fields{client: &testAccountClient{accounts: &expectedAccounts}},
			args:    args{page: listing.Page{Size: 1}},
			wantErr: false,
			errMsg:  "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newService := listing.NewService(test.fields.client)
			accounts, err := newService.GetAccounts(test.args.page)

			if test.wantErr && !strings.Contains(err.Error(), test.errMsg) {
				t.Errorf("GetAccounts() %s: got = %v, want = %v", test.name, err, test.errMsg)
				return
			}
			if !test.wantErr {
				assert.Equal(t, expectedAccounts.Data[0].ID, accounts.Data[0].ID)
				assert.Equal(t, expectedAccounts.Data[0].Type, accounts.Data[0].Type)
			}
		})
	}
}
