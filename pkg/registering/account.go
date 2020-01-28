package registering

// Account consists of data necessary for creating/registering an account.
type Account struct {
	Data `json:"data"`
}

type Data struct {
	Type           string     `json:"type"`
	ID             string     `json:"id"`
	OrganisationID string     `json:"organisation_id"`
	Attributes     Attributes `json:"attributes"`
}

type Attributes struct {
	Country                     string   `json:"country"`
	BaseCurrency                string   `json:"base_currency"`
	AccountNumber               string   `json:"account_number"`
	BankID                      string   `json:"bank_id"`
	BankIDCode                  string   `json:"bank_id_code"`
	Bic                         string   `json:"bic"`
	Iban                        string   `json:"iban"`
	Title                       string   `json:"title"`
	FirstName                   string   `json:"first_name"`
	BankAccountName             string   `json:"bank_account_name"`
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
	AccountClassification       string   `json:"account_classification"`
	JointAccount                bool     `json:"joint_account"`
	AccountMatchingOptOut       bool     `json:"account_matching_opt_out"`
	SecondaryIdentification     string   `json:"secondary_identification"`
}
