package log

import (
	"encoding/json"
	"errors"
	"fmt"
	stdLog "log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const DateTimeFormat = "2006-01-02T15:04:05.999999"

var _ Logger = &Std{}

type Std struct {
	pretty bool
	prefix string
	l      *stdLog.Logger
	level  Level
}

func NewStd(pretty bool, level Level) *Std {
	if level == 0 {
		level = StdDebugLevel
	}
	return &Std{
		pretty: pretty,
		prefix: "",
		l:      stdLog.New(os.Stdout, "", 0),
		level:  level,
	}
}

func (l *Std) Level(level Level) Logger {
	if l == nil {
		return nil
	}

	if l.level == level {
		return l
	}

	return &Std{
		pretty: l.pretty,
		prefix: l.prefix,
		l:      l.l,
		level:  level,
	}
}

func (l *Std) WithField(key string, v any) Logger {
	if l == nil {
		return nil
	}

	prefix, err := createPrefix(key, v)
	if err != nil {
		NewStd(l.pretty, l.level).WithError(err).Errorf("WithField error, passed %#v (type %T)", v, v)
	}

	std := *l
	std.prefix += prefix
	return &std
}

func (l *Std) WithError(err error) Logger {
	if err == nil {
		return l
	}

	log := l.WithField("error", err.Error())
	var stackTracer interface {
		StackTrace() []string
	}
	if errors.As(err, &stackTracer) {
		log = log.WithField("stack", stackTracer.StackTrace())
	}
	return log
}

func (l *Std) Printf(format string, v ...any) {
	_ = l.Logf(StdPrintLevel, format, v...)
}

func (l *Std) Debugf(format string, v ...any) {
	_ = l.Logf(StdDebugLevel, format, v...)
}

func (l *Std) Infof(format string, v ...any) {
	_ = l.Logf(StdInfoLevel, format, v...)
}

func (l *Std) Warnf(format string, v ...any) {
	_ = l.Logf(StdWarnLevel, format, v...)
}

func (l *Std) Errorf(format string, v ...any) {
	_ = l.Logf(StdErrorLevel, format, v...)
}

func (l *Std) Fatalf(format string, v ...any) {
	err := l.Logf(StdFatalLevel, format, v...)
	if err != nil {
		stdLog.Fatalf("level: %s, err: %v", StdFatalLevel, err)
	}
}

func (l *Std) Panicf(format string, v ...any) {
	err := l.Logf(StdPanicLevel, format, v...)
	if err != nil {
		stdLog.Panicf("level: %s, err: %v", StdPanicLevel, err)
	}
}

func (l *Std) Neverf(format string, v ...any) {
	_ = l.Logf(StdNeverLevel, format, v...)
}

// Logf check and print log. Return message and error.
func (l *Std) Logf(lvl Level, format string, v ...any) error {
	if l == nil {
		return fmt.Errorf("nil log.Std: "+format, v...)
	}
	if lvl != StdPrintLevel && l.level > lvl {
		return nil
	}

	var msg strings.Builder
	// Grow underlying bytes buffer as optimization
	msg.Grow(7 + 24 + 10 + (len(l.prefix) + 2) + len(format))
	// Add level
	if l.pretty {
		msg.WriteString(lvl.color().String())
	}
	msg.WriteByte('[')
	msg.WriteString(strings.ToUpper(lvl.String()))
	msg.WriteByte(']')
	if l.pretty {
		msg.WriteString(reset.String())
	}
	// Add date and time
	msg.WriteByte('[')
	msg.WriteString(time.Now().UTC().Format(DateTimeFormat))
	msg.WriteByte(']')
	// Add caller
	if _, file, line, ok := runtime.Caller(2); ok {
		if cwd, err := os.Getwd(); err == nil {
			if rel, err := filepath.Rel(cwd, file); err == nil {
				file = rel
			}
		}
		msg.WriteByte('[')
		msg.WriteString(file)
		msg.WriteByte(':')
		msg.WriteString(strconv.Itoa(line))
		msg.WriteByte(']')
	}
	msg.WriteByte(' ')
	// Add objects as prefix
	if l.prefix != "" {
		msg.WriteString(l.prefix)
		msg.WriteByte(' ')
	}
	// Add message
	fmt.Fprintf(&msg, format, v...)

	// print
	switch lvl {
	case StdFatalLevel:
		l.l.Fatal(msg.String())
	case StdPanicLevel:
		l.l.Panic(msg.String())
	default:
		l.l.Print(msg.String())
	}
	return nil
}

// Level Logger level type.
type Level int

const (
	StdPrintLevel Level = iota
	StdDebugLevel
	StdInfoLevel
	StdWarnLevel
	StdErrorLevel
	StdNeverLevel
	StdFatalLevel
	StdPanicLevel
)

var levels = map[Level]string{
	StdPrintLevel: "print",
	StdDebugLevel: "debug",
	StdInfoLevel:  "info",
	StdWarnLevel:  "warn",
	StdErrorLevel: "error",
	StdFatalLevel: "fatal",
	StdPanicLevel: "panic",
	StdNeverLevel: "never",
}

func (l Level) String() string {
	return levels[l]
}

func (l Level) color() color {
	switch l {
	case StdPrintLevel, StdDebugLevel:
		return blue
	case StdInfoLevel:
		return green
	case StdWarnLevel:
		return yellow
	case StdErrorLevel, StdPanicLevel, StdFatalLevel:
		return red
	default:
		return reset
	}
}

func ParseLevel(str string) (Level, error) {
	for level, levelStr := range levels {
		if str == levelStr {
			return level, nil
		}
	}
	return 0, fmt.Errorf("unknown log level: %s", str)
}

func createPrefix(key string, value any) (string, error) {
	switch val := value.(type) {
	case []byte:
		value = string(val)
	default:
		switch reflect.ValueOf(value).Kind() {
		case reflect.Pointer, reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			b, err := json.Marshal(value)
			if err != nil {
				return "", err
			}
			value = string(b)
		}
	}

	return fmt.Sprintf(`[%s: "%v"]`, key, value), nil
}

type color string

const (
	reset  color = "\033[0m"
	red    color = "\033[31m"
	green  color = "\033[32m"
	yellow color = "\033[33m"
	blue   color = "\033[34m"
)

func (c color) String() string {
	return string(c)
}
