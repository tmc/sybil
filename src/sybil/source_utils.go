package sybil

import (
	"runtime"
	"strings"
)

func getCallerName(skip int) string {
	fpcs := make([]uintptr, 1)
	n := runtime.Callers(skip, fpcs)
	if n == 0 {
		return "default"
	}
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "default"
	}
	parts := strings.Split(fun.Name(), "/")
	return parts[len(parts)-1]
}
