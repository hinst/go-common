package common

import (
	"fmt"
	"runtime/debug"
)

type Exception struct {
	message    string
	stackTrace string
	cause      error
}

func CreateException(message string, cause error) (result *Exception) {
	result = &Exception{
		message:    message,
		stackTrace: string(debug.Stack()),
		cause:      cause,
	}
	return
}

func CreateExceptionIf(message string, cause error) (result *Exception) {
	if cause != nil {
		result = CreateException(message, cause)
	}
	return
}

func (exception *Exception) Error() string {
	var result = exception.message
	if exception.cause != nil {
		result += "\n• caused by: " + exception.cause.Error()
	}
	return result
}

func (exception *Exception) String() string {
	return exception.Error()
}

var _ error = &Exception{}
var _ fmt.Stringer = &Exception{}

func AssertError(errors ...error) {
	for _, e := range errors {
		if !isNil(e) {
			panic(e)
		}
	}
}
