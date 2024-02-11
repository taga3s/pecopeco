package error

type Error struct {
	message string
}

func (e *Error) Error() string {
	return e.message
}

func NewError(s string) *Error {
	return &Error{
		message: s,
	}
}

var NotFoundErr = NewError("not found")
