package consts

// Global constants
const (
	QBO_ITEM_TYPE           = "Service"
	QBO_ACCOUNT_TYPE        = "Income"
	QBO_TXN_TYPE            = "Invoice"
	QBO_INVOICE_DETAIL_TYPE = "SalesItemLineDetail"
	CONTENT_TYPE_XML        = "text/xml"
)

// verb : http verbs
type verb struct {
	Get  string
	Post string
}

// Verb supported http verbs
var (
	Verb = verb{
		Get:  "GET",
		Post: "POST",
	}
)

// QBQueryTypes consts for search type
var (
	QBQueryItemType     = "item"
	QBQueryAccountType  = "account"
	QBQueryCustomerType = "customer"
	// this map makes it easy for validation
	QBQueryTypes = map[string]bool{
		QBQueryItemType:     true,
		QBQueryAccountType:  true,
		QBQueryCustomerType: true,
	}
)
