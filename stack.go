package errors

import (
	"iter"
	"runtime"
)

const maxDepth = 20

type StackTracer interface {
	StackTrace() iter.Seq[runtime.Frame]
}

type stack []uintptr

func newStack() stack {
	var pcs [maxDepth]uintptr
	n := runtime.Callers(3, pcs[:])
	return pcs[:n]
}

func (s stack) StackTrace() iter.Seq[runtime.Frame] {
	return func(yield func(runtime.Frame) bool) {
		frames := runtime.CallersFrames(s)
		for more := true; more; {
			var f runtime.Frame
			f, more = frames.Next()
			if !yield(f) {
				return
			}
		}
	}
}
