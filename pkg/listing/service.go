package listing

import "errors"

const (
	accountIDNotSpecifiedErrMsg    = "failed to specify valid account id"
	invalidPageSizeSpecifiedErrMsg = "failed to provide valid page size, (page size should be greater than 0)"
)

// AccountApiClient interface consists of Fetch and Get methods which fetches an account / lists multiple accounts an account.
type AccountApiClient interface {
	Fetch(id string) (*Account, error)
	List(page Page) (*Accounts, error)
}

type Service struct {
	httpClient AccountApiClient
}

// NewService creates a  listing service with the necessary dependencies
func NewService(client AccountApiClient) *Service {
	if client == nil {
		return nil
	}
	return &Service{
		httpClient: client,
	}
}

func (s *Service) GetAccount(id string) (*Account, error) {
	if id == "" {
		return nil, errors.New(accountIDNotSpecifiedErrMsg)
	}
	account, err := s.httpClient.Fetch(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *Service) GetAccounts(page Page) (*Accounts, error) {
	if page.Size <= 0 {
		return nil, errors.New(invalidPageSizeSpecifiedErrMsg)
	}
	accounts, err := s.httpClient.List(page)
	if err != nil {
		return nil, err
	}
	return accounts, nil
}
