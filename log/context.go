package log

import (
	"context"
)

const (
	TraceIDField = "trace_id"
)

type (
	loggerKey struct{}
)

// NewContext 将Logger存储到context中，并且Logger的字段中存储追踪信息
func NewContext(ctx context.Context, l *Logger, traceID string) context.Context {
	return context.WithValue(ctx, loggerKey{}, l.With(String(TraceIDField, traceID)))
}

// FromContext 从context中，获取日志记录器。当不存在时，返回defaultLogger
func FromContext(ctx context.Context) *Logger {
	if ctx == nil {
		return defaultLogger
	}

	l, ok := ctx.Value(loggerKey{}).(*Logger)
	if !ok {
		return defaultLogger // default logger
	}
	return l
}
