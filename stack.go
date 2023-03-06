package errors

import (
	"fmt"
	"runtime"
)

const depth = 20

type StackTracer interface {
	StackTrace() []string
}

type stack []uintptr

func newStack() stack {
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[:n]
}

func (s stack) StackTrace() []string {
	var trace []string
	frames := runtime.CallersFrames(s)
	for more := true; more; {
		var f runtime.Frame
		f, more = frames.Next()
		trace = append(trace, fmt.Sprintf("%s: %s:%d", f.File, f.Function, f.Line))
	}
	return trace
}
