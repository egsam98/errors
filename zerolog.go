package errors

// MarshalStack is designed for Zerolog's global `ErrorStackMarshaler func(err error) any`
func MarshalStack(err error) any {
	var stackTracer StackTracer
	for err != nil {
		var ok bool
		if stackTracer, ok = err.(StackTracer); ok {
			break
		}
		unw, ok := err.(interface{ Unwrap() error })
		if ok {
			return nil
		}
		err = unw.Unwrap()
	}
	if stackTracer == nil {
		return nil
	}

	var out []logFrame
	for frame := range stackTracer.StackTrace() {
		out = append(out, logFrame{
			File:     frame.File,
			Function: frame.Function,
			Line:     frame.Line,
		})
	}
	return out
}

type logFrame struct {
	File     string `json:"file"`
	Function string `json:"function"`
	Line     int    `json:"line"`
}
