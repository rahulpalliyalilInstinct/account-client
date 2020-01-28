package registering

import (
	"errors"
)

const (
	countryCodeNotSpecifiedErrMsg = "invalid country code specified"

	accountTypeName     = "accounts"
	personalAccountType = "Personal"
)

// UniqueIdentifier interface consis of uuid generator method.
type UniqueIdentifier interface {
	Generator() string
}

// AccountAPIClient interface consists of Create method which creates/registers an account.
type AccountAPIClient interface {
	Create(account Account) error
}

type Service struct {
	httpClient AccountAPIClient
	identifier UniqueIdentifier
}

// NewService creates a  registering service with the necessary dependencies
func NewService(client AccountAPIClient, identifier UniqueIdentifier) *Service {
	if client == nil || identifier == nil {
		return nil
	}
	return &Service{
		httpClient: client,
		identifier: identifier,
	}
}

// CreateAccount creates/registers a given account.
// If the specified account has empty country code, then
// an error is returned.If the organisation id is not
// specified then a new organisation id(uuid) is created
// and will be used to create an account. If the account classification
// is not specified, then `Personal` type is taken by default.
// The account type name is specified as accounts.
func (s *Service) CreateAccount(account Account) error {
	if account.Attributes.Country == "" {
		return errors.New(countryCodeNotSpecifiedErrMsg)
	}

	if account.OrganisationID == "" {
		account.OrganisationID = s.identifier.Generator()
	}

	if account.Attributes.AccountClassification == "" {
		account.Attributes.AccountClassification = personalAccountType
	}

	account.ID = s.identifier.Generator()
	account.Type = accountTypeName

	if err := s.httpClient.Create(account); err != nil {
		return err
	}

	return nil
}
