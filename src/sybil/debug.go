package sybil

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

// extracted from and influenced by
// https://groups.google.com/forum/#!topic/golang-nuts/ct99dtK2Jo4
// use env variable DEBUG=1 to turn on debug output
var ENV_FLAG = os.Getenv("DEBUG")

func Print(args ...interface{}) {
	fmt.Println(args...)
}

func Warn(args ...interface{}) {
	fmt.Fprintln(os.Stderr, append([]interface{}{"Warning:"}, args...)...)
}

func Debug(args ...interface{}) {
	if *FLAGS.DEBUG || ENV_FLAG != "" {
		log.Println(append([]interface{}{callerName(1)}, args...)...)
	}
}

func Error(args ...interface{}) {
	log.Fatalln(append([]interface{}{"ERROR"}, args...)...)
}

// callerName gives the function name (qualified with a package path)
// for the caller after skip frames (where 0 means the current function).
func callerName(skip int) string {
	// Make room for the skip PC.
	var pc [2]uintptr
	n := runtime.Callers(skip+2, pc[:]) // skip + runtime.Callers + callerName
	if n == 0 {
		panic("testing: zero callers found")
	}
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}
