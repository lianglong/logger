package logger

import (
	"context"
	"fmt"
	"log"
	"path"
	"runtime"
	"sync"
)

var (
	driversMu sync.RWMutex
	drivers   = make(map[string]Constructor)
	funcCache sync.Map // map[uintptr]string

	defaultLogger Logger = &noopLogger{} // 默认空实现，避免 panic
)

// Logger 定义日志接口，兼容多种 Web 框架
type Logger interface {
	// 基础日志方法
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	// 分级日志方法
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})

	// Context 集成（框架无关）
	WithContext(ctx context.Context) Logger

	// 结构化日志
	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger

	// 错误处理
	WithError(err error) Logger

	// 刷新缓冲区
	Sync() error
}

// Constructor 驱动构造函数签名
type Constructor func(cfg Config) (Logger, error)

// Register 注册日志驱动（由驱动包的 init() 调用）
func Register(name string, c Constructor) {
	driversMu.Lock()
	defer driversMu.Unlock()
	if _, exists := drivers[name]; exists {
		log.Printf("logger driver %q already registered", name)
		return
	}
	drivers[name] = c
}

// New 创建日志实例
func New(driver string, cfg Config) (Logger, error) {
	driversMu.RLock()
	defer driversMu.RUnlock()
	c, ok := drivers[driver]
	if !ok {
		return nil, fmt.Errorf("logger driver %q not found", driver)
	}
	return c(cfg)
}

// MustNew 创建日志实例，失败则 panic（便捷方法）
func MustNew(driver string, cfg Config) Logger {
	l, err := New(driver, cfg)
	if err != nil {
		log.Fatal(err)
	}
	return l
}

// Drivers 返回已注册的驱动列表
func Drivers() []string {
	driversMu.RLock()
	defer driversMu.RUnlock()
	names := make([]string, 0, len(drivers))
	for name := range drivers {
		names = append(names, name)
	}
	return names
}

// GetCaller 获取调用者信息（供驱动使用）
func GetCaller(skip int) map[string]interface{} {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return nil
	}
	var fn string
	if v, hit := funcCache.Load(pc); hit {
		fn = v.(string)
	} else {
		fn = path.Base(runtime.FuncForPC(pc).Name())
		funcCache.Store(pc, fn)
	}
	return map[string]interface{}{
		"file": file,
		"line": line,
		"func": fn,
	}
}

// noopLogger 空实现，防止 nil panic
type noopLogger struct{}

func (n *noopLogger) Print(...interface{})                                  {}
func (n *noopLogger) Printf(string, ...interface{})                         {}
func (n *noopLogger) Println(...interface{})                                {}
func (n *noopLogger) Debug(...interface{})                                  {}
func (n *noopLogger) Debugf(string, ...interface{})                         {}
func (n *noopLogger) Info(...interface{})                                   {}
func (n *noopLogger) Infof(string, ...interface{})                          {}
func (n *noopLogger) Warn(...interface{})                                   {}
func (n *noopLogger) Warnf(string, ...interface{})                          {}
func (n *noopLogger) Error(...interface{})                                  {}
func (n *noopLogger) Errorf(string, ...interface{})                         {}
func (n *noopLogger) Fatal(...interface{})                                  {}
func (n *noopLogger) Fatalf(string, ...interface{})                         {}
func (n *noopLogger) WithContext(ctx context.Context) Logger                { return n }
func (n *noopLogger) WithField(key string, value interface{}) Logger        { return n }
func (n *noopLogger) WithFields(fields map[string]interface{}) Logger       { return n }
func (n *noopLogger) WithError(err error) Logger                            { return n }
func (n *noopLogger) Sync() error                                           { return nil }
