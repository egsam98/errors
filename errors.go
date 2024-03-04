package errors

import (
	"errors"
	"fmt"
)

type withStack struct {
	error
	stack
}

func (w *withStack) Unwrap() error {
	return w.error
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

func Wrapf(err error, format string, args ...any) error {
	return &withStack{
		error: fmt.Errorf(format+": %w", append(args, err)...),
		stack: getOrNewStack(err),
	}
}

func WrapRightf(err error, format string, args ...any) error {
	return &withStack{
		error: fmt.Errorf("%w: "+format, append([]any{err}, args...)...),
		stack: getOrNewStack(err),
	}
}

func Wrap(err error, message string) error {
	return Wrapf(err, message)
}

func WithStack(err error) error {
	if _, ok := err.(*withStack); ok {
		return err
	}
	return &withStack{
		error: err,
		stack: newStack(),
	}
}

func getOrNewStack(err error) stack {
	if w, ok := err.(*withStack); ok {
		return w.stack
	}
	return newStack()
}
