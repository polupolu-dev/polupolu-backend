package errors

const (
	ErrNotFound     = "resource not found"
	ErrInvalidInput = "invalid input"
	ErrInternal     = "internal error"
)

type Error struct {
	Code    string
	Message string
}

func NewError(code, message string) *Error {
	return &Error{Code: code, Message: message}
}
