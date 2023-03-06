package log

type Noop struct{}

func (n *Noop) Printf(format string, v ...any)   {}
func (n *Noop) Debugf(format string, v ...any)   {}
func (n *Noop) Infof(format string, v ...any)    {}
func (n *Noop) Warnf(format string, v ...any)    {}
func (n *Noop) Errorf(format string, v ...any)   {}
func (n *Noop) Neverf(format string, v ...any)   {}
func (n *Noop) Fatalf(format string, v ...any)   {}
func (n *Noop) Panicf(format string, v ...any)   {}
func (n *Noop) WithField(k string, v any) Logger { return n }
func (n *Noop) WithError(err error) Logger       { return n }
func (n *Noop) Level(level Level) Logger         { return n }
