package deregistering

import "errors"

const (
	accountIDNotSpecifiedErrMsg = "failed to specify valid account id"

	defaultVersion = "0"
)

// AccountApiClient interface consists of Delete method which deletes/deregisters an account.
type AccountApiClient interface {
	Delete(account Account) error
}

type Service struct {
	httpClient AccountApiClient
}

// NewService creates a  deregistering service with the necessary dependencies
func NewService(client AccountApiClient) *Service {
	if client == nil {
		return nil
	}
	return &Service{
		httpClient: client,
	}
}

// DeleteAccount deletes a specified account. If there is no
// id specified then accountIDNotSpecifiedErrMsg is returned
// if version is not specified the defaultVersion(0) is used
// to delete the record.
func (s *Service) DeleteAccount(account Account) error {
	if account.ID == "" {
		return errors.New(accountIDNotSpecifiedErrMsg)
	}

	if account.Version == "" {
		account.Version = defaultVersion
	}

	if err := s.httpClient.Delete(account); err != nil {
		return err
	}

	return nil
}
