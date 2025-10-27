package logger

import "context"

type contextKey string

const (
	loggerKey    contextKey = "logger"
	requestIDKey contextKey = "request_id"
	userIDKey    contextKey = "user_id"
	traceIDKey   contextKey = "trace_id"
)

// WithLogger 将 logger 存入 context
func WithLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, loggerKey, log)
}

// FromContext 从 context 获取 logger（框架无关）
func FromContext(ctx context.Context) Logger {
	if log, ok := ctx.Value(loggerKey).(Logger); ok {
		return log
	}
	return defaultLogger
}

// WithRequestID 设置 request_id
func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

// GetRequestID 获取 request_id
func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(requestIDKey).(string); ok {
		return id
	}
	return ""
}

// WithUserID 设置 user_id
func WithUserID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, userIDKey, id)
}

// GetUserID 获取 user_id
func GetUserID(ctx context.Context) string {
	if id, ok := ctx.Value(userIDKey).(string); ok {
		return id
	}
	return ""
}

// WithTraceID 设置 trace_id（分布式追踪）
func WithTraceID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, traceIDKey, id)
}

// GetTraceID 获取 trace_id
func GetTraceID(ctx context.Context) string {
	if id, ok := ctx.Value(traceIDKey).(string); ok {
		return id
	}
	return ""
}

// WithField 设置自定义字段
func WithField(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, contextKey(key), value)
}

// GetField 获取自定义字段
func GetField(ctx context.Context, key string) interface{} {
	return ctx.Value(contextKey(key))
}
