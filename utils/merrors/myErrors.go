package merrors

const (
	ErrNotFound     = "resource not found"
	ErrInvalidInput = "invalid input"
	ErrInternal     = "internal error"
)

type Error struct {
	Code    string
	Message error
}

func NewError(code string, message error) *Error {
	return &Error{Code: code, Message: message}
}
