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

func New(text string) error {
	return &errorString{text}
}
