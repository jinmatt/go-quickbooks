package consts

// Global constants
const (
	QBFaultType    = "AUTHENTICATION"
	ContextTypeXML = "text/xml"
	QBAccountType  = "Income"
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
