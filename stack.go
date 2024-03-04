package errors

import (
	"runtime"
	"strconv"
	"strings"
)

const maxDepth = 20

type StackTracer interface {
	StackTrace() []string
}

type stack []uintptr

func newStack() stack {
	var pcs [maxDepth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[:n]
}

func (s stack) StackTrace() (trace []string) {
	frames := runtime.CallersFrames(s)
	for more := true; more; {
		var f runtime.Frame
		f, more = frames.Next()

		var str strings.Builder
		str.WriteString(f.File)
		str.WriteString(": ")
		str.WriteString(f.Function)
		str.WriteByte(':')
		str.WriteString(strconv.Itoa(f.Line))

		trace = append(trace, str.String())
	}
	return
}
