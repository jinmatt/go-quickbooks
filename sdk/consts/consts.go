package consts

// Global constants
const (
	QBFaultType         = "AUTHENTICATION"
	ContextTypeXML      = "text/xml"
	QBAccountIncomeType = "Income"
	QBItemServiceType   = "Service"
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
