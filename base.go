package logger

import (
	"runtime"
	"strconv"
)

// FILE ...
// @Description:
// @return string
//
func FILE() string {
	_, file, _, _ := runtime.Caller(2)
	return file
}

// LINE ...
// @Description:
// @return string
//
func LINE() string {
	_, _, line, _ := runtime.Caller(2)
	return strconv.Itoa(line)
}
