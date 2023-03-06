package errors

import (
	"errors"
	"fmt"
)

type withStack struct {
	error
	stack
}

func New(msg string) error {
	return &withStack{
		error: errors.New(msg),
		stack: newStack(),
	}
}

func Errorf(format string, args ...any) error {
	return &withStack{
		error: fmt.Errorf(format, args...),
		stack: newStack(),
	}
}

func Wrap(err error, format string, args ...any) error {
	var st stack
	if w, ok := err.(*withStack); ok {
		st = w.stack
	} else {
		st = newStack()
	}
	return &withStack{
		error: fmt.Errorf(format+": %w", append(args, err)...),
		stack: st,
	}
}

func (w *withStack) Unwrap() error {
	return w.error
}
