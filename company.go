package quickbooks

import (
	"encoding/json"
	"fmt"
)

// CompanyObject the complete quickbooks company object type
type CompanyObject struct {
	CompanyInfo struct {
		CompanyName string `json:"CompanyName"`
		LegalName   string `json:"LegalName"`
		CompanyAddr struct {
			ID                     string `json:"Id"`
			Line1                  string `json:"Line1"`
			City                   string `json:"City"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			PostalCode             string `json:"PostalCode"`
			Lat                    string `json:"Lat"`
			Long                   string `json:"Long"`
		} `json:"CompanyAddr"`
		CustomerCommunicationAddr struct {
			ID                     string `json:"Id"`
			Line1                  string `json:"Line1"`
			City                   string `json:"City"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			PostalCode             string `json:"PostalCode"`
			Lat                    string `json:"Lat"`
			Long                   string `json:"Long"`
		} `json:"CustomerCommunicationAddr"`
		LegalAddr struct {
			ID                     string `json:"Id"`
			Line1                  string `json:"Line1"`
			City                   string `json:"City"`
			CountrySubDivisionCode string `json:"CountrySubDivisionCode"`
			PostalCode             string `json:"PostalCode"`
			Lat                    string `json:"Lat"`
			Long                   string `json:"Long"`
		} `json:"LegalAddr"`
		PrimaryPhone struct {
		} `json:"PrimaryPhone"`
		CompanyStartDate     string `json:"CompanyStartDate"`
		FiscalYearStartMonth string `json:"FiscalYearStartMonth"`
		Country              string `json:"Country"`
		Email                struct {
			Address string `json:"Address"`
		} `json:"Email"`
		WebAddr struct {
		} `json:"WebAddr"`
		SupportedLanguages string `json:"SupportedLanguages"`
		NameValue          []struct {
			Name  string `json:"Name"`
			Value string `json:"Value"`
		} `json:"NameValue"`
		Domain    string `json:"domain"`
		Sparse    bool   `json:"sparse"`
		ID        string `json:"Id"`
		SyncToken string `json:"SyncToken"`
		MetaData  struct {
			CreateTime      string `json:"CreateTime"`
			LastUpdatedTime string `json:"LastUpdatedTime"`
		} `json:"MetaData"`
	} `json:"CompanyInfo"`
	Time string `json:"time"`
}

// GetCompanyInfo returns company info based on realm ID/company ID passed to NewClient options
func (qb *quickbooks) GetCompanyInfo() (*CompanyObject, error) {
	endpoint := fmt.Sprintf("/company/%s/companyinfo/%s", qb.realmID, qb.realmID)

	res, err := qb.makeGetRequest(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	company := CompanyObject{}
	err = json.NewDecoder(res.Body).Decode(&company)
	if err != nil {
		return nil, err
	}

	return &company, nil
}
