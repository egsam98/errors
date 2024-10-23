package errors

import (
	"errors"
	"fmt"
)

func Is(err error, target error) bool { return errors.Is(err, target) }
func As(err error, target any) bool   { return errors.As(err, target) }

func New(msg string) error {
	return &withStack{
		error: errors.New(msg),
		stack: newStack(),
	}
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...any) error {
	return &withStack{
		error: fmt.Errorf(format, args...),
		stack: newStack(),
	}
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the format specifier.
// If err is nil, Wrap returns nil.
func Wrap(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return &withStack{
		error: fmt.Errorf(format+": %w", append(args, err)...),
		stack: getOrNewStack(err),
	}
}

// WrapR behaves similarly to Wrap but appends format and args on the right side of formatted text
func WrapR(err error, format string, args ...any) error {
	if err == nil {
		return nil
	}
	return &withStack{
		error: fmt.Errorf("%w: "+format, append([]any{err}, args...)...),
		stack: getOrNewStack(err),
	}
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	if err == nil {
		return nil
	}
	if _, ok := err.(*withStack); ok { //nolint:errorlint
		return err
	}
	return &withStack{
		error: err,
		stack: newStack(),
	}
}

func getOrNewStack(err error) stack {
	if w, ok := err.(*withStack); ok { //nolint:errorlint
		return w.stack
	}
	return newStack()
}

type withStack struct {
	error
	stack
}

func (w *withStack) Unwrap() error {
	return w.error
}

func (w *withStack) MarshalText() ([]byte, error) {
	return []byte(w.error.Error()), nil
}
