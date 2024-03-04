package errors

import stderr "errors"

func Is(err error, target error) bool { return stderr.Is(err, target) }
func As(err error, target any) bool   { return stderr.As(err, target) }
