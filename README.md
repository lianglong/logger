# Logger

通用的 Go 日志库接口，支持多种底层实现和 Web 框架。

## 特性

- 🔌 **可插拔驱动**：支持 Zap、Logrus、Zerolog 等多种实现
- 🌐 **框架无关**：基于标准 `context.Context`，兼容 Gin、Iris、Echo 等
- 📝 **结构化日志**：支持字段和上下文
- 🎯 **零依赖**：核心包无第三方依赖
- 🚀 **高性能**：函数名缓存、字段提取器优化

## 安装

```bash
# 核心包
go get github.com/lianglong/logger

# Zap 驱动（可选）
go get github.com/lianglong/logger-zap