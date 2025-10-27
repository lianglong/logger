package logger

import "context"

// FieldExtractor 从 context 中提取字段的接口
type FieldExtractor interface {
	Extract(ctx context.Context) map[string]interface{}
}

// FieldExtractorFunc 函数适配器
type FieldExtractorFunc func(ctx context.Context) map[string]interface{}

func (f FieldExtractorFunc) Extract(ctx context.Context) map[string]interface{} {
	return f(ctx)
}

// 常用字段名常量
const (
	FieldRequestID = "request_id"
	FieldUserID    = "user_id"
	FieldTraceID   = "trace_id"
	FieldError     = "error"
)

// DefaultExtractors 默认字段提取器
var DefaultExtractors = []FieldExtractor{
	RequestIDExtractor,
	UserIDExtractor,
	TraceIDExtractor,
}

// RequestIDExtractor 提取 request_id
var RequestIDExtractor = FieldExtractorFunc(func(ctx context.Context) map[string]interface{} {
	if id := GetRequestID(ctx); id != "" {
		return map[string]interface{}{FieldRequestID: id}
	}
	return nil
})

// UserIDExtractor 提取 user_id
var UserIDExtractor = FieldExtractorFunc(func(ctx context.Context) map[string]interface{} {
	if id := GetUserID(ctx); id != "" {
		return map[string]interface{}{FieldUserID: id}
	}
	return nil
})

// TraceIDExtractor 提取 trace_id
var TraceIDExtractor = FieldExtractorFunc(func(ctx context.Context) map[string]interface{} {
	if id := GetTraceID(ctx); id != "" {
		return map[string]interface{}{FieldTraceID: id}
	}
	return nil
})

// ExtractFields 使用提取器列表从 context 提取所有字段
func ExtractFields(ctx context.Context, extractors []FieldExtractor) map[string]interface{} {
	if len(extractors) == 0 {
		extractors = DefaultExtractors
	}

	fields := make(map[string]interface{})
	for _, extractor := range extractors {
		if extracted := extractor.Extract(ctx); extracted != nil {
			for k, v := range extracted {
				fields[k] = v
			}
		}
	}
	return fields
}
