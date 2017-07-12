package quickbooks

import (
	"encoding/json"
	"fmt"

	"github.com/jinmatt/go-quickbooks/sdk/types"
)

// GetCompanyInfo returns company info based on realm ID/company ID passed to NewClient options
func (qb *quickbooks) GetCompanyInfo() (*types.Company, error) {
	endpoint := fmt.Sprintf("/company/%s/companyinfo/%s", qb.realmID, qb.realmID)

	res, err := qb.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	company := types.Company{}
	err = json.NewDecoder(res.Body).Decode(&company)
	if err != nil {
		return nil, err
	}

	return &company, nil
}
