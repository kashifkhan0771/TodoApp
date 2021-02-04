package errors

// const will be used for the status of api handlers
const (
	NotFound            = "Not_Found"
	BadRequest          = "Bad_Request"
	NoContent           = "No_Content"
	InternalServerError = "Internal_Server_Error"
	Timeout             = "Timeout"
)

// APIError struct contains the code and message of error
type APIError struct {
	code    string
	message string
}

// Error returns error message.
func (a *APIError) Error() string {
	return a.message
}

// IsError returns error type.
func (a *APIError) IsError(errType string) bool {
	return a.code == errType
}

// NewAPIError returns api error.
func NewAPIError(errType, message string) *APIError {
	return &APIError{
		code:    errType,
		message: message,
	}
}
