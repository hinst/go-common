package common

import "fmt"

type Exception struct {
	message string
	cause   error
}

func CreateException(message string, cause error) (result Exception) {
	result.message = message
	result.cause = cause
	return
}

func (exception Exception) Error() string {
	var result = exception.message
	if exception.cause != nil {
		result += "\n" + exception.cause.Error()
	}
	return result
}

func (exception Exception) String() string {
	return exception.Error()
}

var _ error = Exception{}
var _ fmt.Stringer = Exception{}

func AssertWrapped(e error, message string) {
	if e != nil {
		if len(message) > 0 {
			panic(Exception{message: message, cause: e})
		} else {
			panic(e)
		}
	}
}

func AssertError(errors ...error) {
	for _, e := range errors {
		if nil != e {
			panic(e)
		}
	}
}
