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

// AccountApiClient interface consists of Create method which creates/registers an account.
type AccountApiClient interface {
	Create(account Account) error
}

type Service struct {
	httpClient AccountApiClient
	identifier UniqueIdentifier
}

// NewService creates a  registering service with the necessary dependencies
func NewService(client AccountApiClient, identifier UniqueIdentifier) *Service {
	if client == nil || identifier == nil {
		return nil
	}
	return &Service{
		httpClient: client,
		identifier: identifier,
	}
}

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
