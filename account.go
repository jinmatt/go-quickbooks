package quickbooks

import (
	"encoding/json"
	"fmt"
)

// AccountObject the complete quickbooks account object type
type AccountObject struct {
	Account *Account `json:"Account"`
	Time    string   `json:"time"`
}

// Account quickbooks account type
type Account struct {
	ID                            string `json:"Id,omitempty"`
	Name                          string `json:"Name"`
	SubAccount                    bool   `json:"SubAccount,omitempty"`
	FullyQualifiedName            string `json:"FullyQualifiedName,omitempty"`
	Active                        bool   `json:"Active,omitempty"`
	Classification                string `json:"Classification,omitempty"`
	AccountType                   string `json:"AccountType"`
	AccountSubType                string `json:"AccountSubType,omitempty"`
	CurrentBalance                int    `json:"CurrentBalance,omitempty"`
	CurrentBalanceWithSubAccounts int    `json:"CurrentBalanceWithSubAccounts,omitempty"`
	CurrencyRef                   *struct {
		Value string `json:"value"`
		Name  string `json:"name"`
	} `json:"CurrencyRef,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Sparse    bool   `json:"sparse,omitempty"`
	SyncToken string `json:"SyncToken,omitempty"`
	MetaData  *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
}

// CreateAccount creates a chart of account on quickbooks
func (qb *quickbooks) CreateAccount(account Account) (*AccountObject, error) {
	endpoint := fmt.Sprintf("/company/%s/account", qb.realmID)

	res, err := qb.makePostRequest(endpoint, account)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newAccount := AccountObject{}
	err = json.NewDecoder(res.Body).Decode(&newAccount)
	if err != nil {
		return nil, err
	}

	return &newAccount, nil
}
