package consts

// Global constants
const (
	QBItemType          = "Service"
	QBAccountType       = "Income"
	QBTxnType           = "Invoice"
	QBInvoiceDetailType = "SalesItemLineDetail"
	QBFaultType         = "AUTHENTICATION"
	ContextTypeXML      = "text/xml"
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
