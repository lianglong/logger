package logger

import "io"

// Config 日志配置
type Config struct {
	// 日志级别
	Level Level

	// 输出目标（stdout/stderr/文件等）
	Output io.Writer

	// 时间格式（如: "2006-01-02 15:04:05"）
	TimeLayout string

	// 是否记录调用者信息（文件名、行号）
	WithCaller bool

	// 字段提取器列表（从 context 提取字段）
	FieldExtractors []FieldExtractor

	// 驱动专属配置（可选）
	Extra map[string]interface{}
}

// Validate 验证配置
func (c *Config) Validate() error {
	if c.Output == nil {
		return ErrInvalidConfig("output is required")
	}
	return nil
}

// ErrInvalidConfig 配置错误
type ErrInvalidConfig string

func (e ErrInvalidConfig) Error() string {
	return "invalid config: " + string(e)
}
