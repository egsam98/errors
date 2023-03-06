package log

// Logger describe methods with different log levels.
// Suffix "f" means formatted message.
type Logger interface {
	Printf(format string, v ...any)
	Debugf(format string, v ...any)
	Infof(format string, v ...any)
	Warnf(format string, v ...any)
	Errorf(format string, v ...any)
	Fatalf(format string, v ...any)
	Panicf(format string, v ...any)
	Neverf(format string, v ...any)
	WithField(k string, v any) Logger
	WithError(err error) Logger
	Level(level Level) Logger
}
