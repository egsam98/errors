package log

import (
	"context"
	"sync"
)

type contextKeyT string

var (
	contextKey           = contextKeyT("gitlab.com/bro_ag/go-lib/service/log")
	defaultLogger Logger = new(Noop)
	mx            sync.RWMutex
)

// FromContext returns the Logger bound to the context or defaultLogger if it not exists.
func FromContext(ctx context.Context) Logger {
	if l, ok := ctx.Value(contextKey).(Logger); ok {
		return l
	}
	mx.RLock()
	defer mx.RUnlock()
	return defaultLogger
}

// NewContext returns a copy of the parent context
// and associates it with a Logger.
func NewContext(ctx context.Context, l Logger) context.Context {
	return context.WithValue(ctx, contextKey, l)
}

func SetDefault(logger Logger) {
	mx.Lock()
	defer mx.Unlock()
	defaultLogger = logger
}
