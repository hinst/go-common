package common

import (
	"os"
	"path"
	"unsafe"
)

var SizeOfInt = unsafe.Sizeof(int(0))
var BitSizeOfInt = SizeOfInt * 8

var SizeOfUint = unsafe.Sizeof(uint(0))
var BitSizeOfUInt = SizeOfUint * 8

func GetExecutableDir() string {
	var filePath, e = os.Executable()
	AssertError(CreateExceptionIf("Unable to get executable file path", e))
	return path.Dir(filePath)
}