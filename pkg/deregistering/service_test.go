package deregistering_test

import (
	"errors"
	"strings"
	"testing"

	"account-client/pkg/deregistering"
	"github.com/stretchr/testify/assert"
)

type testAccountClient struct {
	err error
}

func (t *testAccountClient) Delete(account deregistering.Account) error {
	return t.err
}

func TestNewService(t *testing.T) {
	t.Parallel()
	type fields struct {
		client deregistering.AccountApiClient
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
			name:          "returns valid service",
			fields:        fields{client: &testAccountClient{}},
			invalidHandle: false,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newService := deregistering.NewService(test.fields.client)
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
		client deregistering.AccountApiClient
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
			name:    "empty account id specified",
			args:    args{id: ""},
			wantErr: true,
			errMsg:  "failed to specify valid account id",
		},
		{
			name:    "failed to deregister account",
			args:    args{id: "valid id"},
			fields:  fields{client: &testAccountClient{err: errors.New(expectedErrMsg)}},
			wantErr: true,
			errMsg:  expectedErrMsg,
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newService := deregistering.NewService(test.fields.client)
			err := newService.DeleteAccount(deregistering.Account{
				ID: test.args.id,
			})

			if test.wantErr && !strings.Contains(err.Error(), test.errMsg) {
				t.Errorf("DeleteAccount() %s: got = %v, want = %v", test.name, err, test.errMsg)
				return
			}
			if !test.wantErr {
				assert.Nil(t, err)
			}
		})
	}

}
