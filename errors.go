package common

func WrapError(message string, e error) *ErrorWrapper {
	if e != nil {
		return &ErrorWrapper{Message: message, InnerError: e}
	} else {
		return nil
	}
}

type ErrorWrapper struct {
	Message    string
	InnerError error
}

var _ error = &ErrorWrapper{}

func (e *ErrorWrapper) Error() string {
	return e.Message + " CAUSED BY\n" + e.InnerError.Error()
}

func AssertError(errors ...error) {
	for _, e := range errors {
		if nil != e {
			panic(e)
		}
	}
}
