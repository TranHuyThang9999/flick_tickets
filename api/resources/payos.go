package resources

type PayOSError struct {
	code    string
	message string
}

func NewPayOSError(code, message string) *PayOSError {
	return &PayOSError{
		code:    code,
		message: message,
	}
}

func (e *PayOSError) Error() string {
	return e.message
}
