package errors

// Upstream errors
var (
	QBApiFailure = NewSDKError("upstream-1", "Quickbooks API responded with error")
)
