package errors

// Quickbook errors
var (
	QBApiFailure  = NewSDKError("Q101", "Quickbooks API responded with error")
	QBAuthFailure = NewSDKError("Q102", "Quickbooks authentication failed")
)
