package errors

// SDKError error type
type SDKError struct {
	errorCode    string
	errorMessage string
}

// NewSDKError creates a new error of type *SDKError
func NewSDKError(errorCode, errorMessage string) *SDKError {
	return &SDKError{errorCode, errorMessage}
}

func (e *SDKError) Error() string {
	return "Error: " + e.errorCode + " Message: " + e.errorMessage
}

// ErrorCode returns error code from SDKError type
func (e *SDKError) ErrorCode() string {
	return e.errorCode
}

// ErrorMessage returns error message from SDKError type
func (e *SDKError) ErrorMessage() string {
	return e.errorMessage
}
