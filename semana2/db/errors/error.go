package errors

type error interface {
	Error() string
}

type errorString struct {
	str string
}

func (e *errorString) Error() string {
	return e.str
}
