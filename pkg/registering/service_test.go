package registering_test

import (
	"errors"
	"strings"
	"testing"

	"account-client/pkg/registering"

	"github.com/stretchr/testify/assert"
)

type testAccountClient struct {
	err error
}

type testUniqueIdentifier struct {
	text string
}

func (t *testAccountClient) Create(account registering.Account) error {
	return t.err
}

func (t *testUniqueIdentifier) Generator() string {
	return t.text
}

func TestNewService(t *testing.T) {
	t.Parallel()
	type fields struct {
		client     registering.AccountAPIClient
		identifier registering.UniqueIdentifier
	}
	tests := []struct {
		name          string
		fields        fields
		invalidHandle bool
	}{
		{
			name: "returns nil service handle when a nil account client is passed",
			fields: fields{
				identifier: &testUniqueIdentifier{},
			},
			invalidHandle: true,
		},
		{
			name: "returns nil service handle when a nil identifier is passed",
			fields: fields{
				client: &testAccountClient{},
			},
			invalidHandle: true,
		},
		{
			name: "returns valid service",
			fields: fields{
				client:     &testAccountClient{},
				identifier: &testUniqueIdentifier{},
			},
			invalidHandle: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newService := registering.NewService(test.fields.client, test.fields.identifier)
			if test.invalidHandle {
				assert.Nil(t, newService)
				return
			}
			assert.NotNil(t, newService)
		})
	}
}

func TestService_CreateAccount(t *testing.T) {
	t.Parallel()
	expectedErrMsg := "something bad happened"
	type fields struct {
		client     registering.AccountAPIClient
		identifier registering.UniqueIdentifier
	}
	type args struct {
		countryCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		errMsg  string
	}{
		{
			name:    "invalid country code specified",
			args:    args{countryCode: ""},
			wantErr: true,
			errMsg:  "invalid country code",
		},
		{
			name: "failed to create account",
			fields: fields{
				identifier: &testUniqueIdentifier{text: "some ulid"},
				client:     &testAccountClient{err: errors.New(expectedErrMsg)},
			},
			args:    args{countryCode: "valid country code"},
			wantErr: true,
			errMsg:  expectedErrMsg,
		},
		{
			name: "successfully create account",
			fields: fields{
				identifier: &testUniqueIdentifier{text: "some ulid"},
				client:     &testAccountClient{},
			},
			args:    args{countryCode: "valid country code"},
			wantErr: false,
			errMsg:  "",
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newService := registering.NewService(test.fields.client, test.fields.identifier)
			err := newService.CreateAccount(registering.Account{
				Data: registering.Data{
					Attributes: registering.Attributes{Country: test.args.countryCode},
				},
			})

			if test.wantErr && !strings.Contains(err.Error(), test.errMsg) {
				t.Errorf("CreateAccount() %s: got = %v, want = %v", test.name, err, test.errMsg)
				return
			}
			if !test.wantErr {
				assert.Nil(t, err)
			}
		})
	}
}
