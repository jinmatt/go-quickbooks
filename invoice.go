package quickbooks

import (
	"encoding/json"
	"fmt"
)

// InvoiceObject the complete quickbooks invoice object type
type InvoiceObject struct {
	Invoice *Invoice `json:"Invoice"`
	Time    string   `json:"time"`
}

// Invoice quickbooks invoice type
type Invoice struct {
	ID        string `json:"Id,omitempty"`
	Deposit   int    `json:"Deposit,omitempty"`
	Domain    string `json:"domain,omitempty"`
	Sparse    bool   `json:"sparse,omitempty"`
	SyncToken string `json:"SyncToken,omitempty"`
	MetaData  *struct {
		CreateTime      string `json:"CreateTime"`
		LastUpdatedTime string `json:"LastUpdatedTime"`
	} `json:"MetaData,omitempty"`
	CustomField *[]struct {
		DefinitionID string `json:"DefinitionId"`
		Name         string `json:"Name"`
		Type         string `json:"Type"`
		StringValue  string `json:"StringValue"`
	} `json:"CustomField,omitempty"`
	DocNumber string `json:"DocNumber,omitempty"`
	TxnDate   string `json:"TxnDate,omitempty"`
	LinkedTxn *[]struct {
		TxnID   string `json:"TxnId"`
		TxnType string `json:"TxnType"`
	} `json:"LinkedTxn,omitempty"`
	Line         []InvoiceLine `json:"Line"`
	TxnTaxDetail *struct {
		TxnTaxCodeRef *struct {
			Value string `json:"value"`
		} `json:"TxnTaxCodeRef,omitempty"`
		TotalTax float64 `json:"TotalTax,omitempty"`
		TaxLine  *[]struct {
			Amount        float64 `json:"Amount"`
			DetailType    string  `json:"DetailType"`
			TaxLineDetail *struct {
				TaxRateRef *struct {
					Value string `json:"value"`
				} `json:"TaxRateRef,omitempty"`
				PercentBased     bool    `json:"PercentBased"`
				TaxPercent       int     `json:"TaxPercent"`
				NetAmountTaxable float64 `json:"NetAmountTaxable"`
			} `json:"TaxLineDetail,omitempty"`
		} `json:"TaxLine,omitempty"`
	} `json:"TxnTaxDetail,omitempty"`
	CustomerRef  *CustomerRef `json:"CustomerRef"`
	CustomerMemo *struct {
		Value string `json:"value"`
	} `json:"CustomerMemo,omitempty"`
	BillAddr     *Address `json:"BillAddr,omitempty"`
	ShipAddr     *Address `json:"ShipAddr,omitempty"`
	SalesTermRef *struct {
		Value string `json:"value"`
	} `json:"SalesTermRef,omitempty"`
	DueDate               string            `json:"DueDate,omitempty"`
	TotalAmt              float64           `json:"TotalAmt,omitempty"`
	ApplyTaxAfterDiscount bool              `json:"ApplyTaxAfterDiscount,omitempty"`
	PrintStatus           string            `json:"PrintStatus,omitempty"`
	EmailStatus           string            `json:"EmailStatus,omitempty"`
	BillEmail             *PrimaryEmailAddr `json:"BillEmail,omitempty"`
	Balance               float64           `json:"Balance,omitempty"`
}

type CustomerRef struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}

// InvoiceLine quickbooks invoice line item type
type InvoiceLine struct {
	ID                  string               `json:"Id,omitempty"`
	LineNum             int                  `json:"LineNum,omitempty"`
	Description         string               `json:"Description,omitempty"`
	Amount              float64              `json:"Amount"`
	DetailType          string               `json:"DetailType"`
	SalesItemLineDetail *SalesItemLineDetail `json:"SalesItemLineDetail"`
	SubTotalLineDetail  *struct {
	} `json:"SubTotalLineDetail,omitempty"`
}

// SalesItemLineDetail invoice sales line item detail type
type SalesItemLineDetail struct {
	ItemRef    *ItemRef    `json:"ItemRef"`
	UnitPrice  int         `json:"UnitPrice,omitempty"`
	Qty        int         `json:"Qty,omitempty"`
	TaxCodeRef *TaxCodeRef `json:"TaxCodeRef,omitempty"`
}

type ItemRef struct {
	Value string `json:"value"`
	Name  string `json:"name,omitempty"`
}

type TaxCodeRef struct {
	Value string `json:"value"`
}

// CreateInvoice creates an invoice on quickbooks
func (qb *quickbooks) CreateInvoice(invoice Invoice) (*InvoiceObject, error) {
	endpoint := fmt.Sprintf("/company/%s/invoice", qb.realmID)

	res, err := qb.makePostRequest(endpoint, invoice)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	newInvoice := InvoiceObject{}
	err = json.NewDecoder(res.Body).Decode(&newInvoice)
	if err != nil {
		return nil, err
	}

	return &newInvoice, nil
}
