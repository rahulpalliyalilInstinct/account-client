package listing

import "time"

// Accounts consists of a collection of data and links.
// Links will have next and previous links to pages.
type Accounts struct {
	Data  []Data     `json:"data"`
	Links Navigation `json:"links"`
}

type Account struct {
	Data Data `json:"data"`
}

type Data struct {
	CreatedOn      time.Time  `json:"created_on"`
	ID             string     `json:"id"`
	ModifiedOn     time.Time  `json:"modified_on"`
	OrganisationID string     `json:"organisation_id"`
	Type           string     `json:"type"`
	Version        int        `json:"version"`
	Attributes     Attributes `json:"attributes"`
}

type Attributes struct {
	AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
	AccountClassification       string   `json:"account_classification"`
	JointAccount                bool     `json:"joint_account"`
	AccountMatchingOptOut       bool     `json:"account_matching_opt_out"`
	SecondaryIdentification     string   `json:"secondary_identification"`
}

type Navigation struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
