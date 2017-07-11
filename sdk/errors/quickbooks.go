package errors

// Quickbook errors
var (
	QBApiFailure  = NewSDKError("quickbooks-1", "Quickbooks API responded with error")
	QBAuthFailure = NewSDKError("quickbooks-2", "Quickbooks authentication failed")
)
