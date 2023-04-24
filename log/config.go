package log

import (
	"fmt"
	goKit "practice.com/go-kit"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config 日志相关配置
type Config struct {
	// Dir 日志输出目录
	Dir string
	// Name 日志文件名称
	Name string
	// Level 日志初始等级
	Level string
	// 日志初始化字段
	Fields []zap.Field
	// 是否添加调用者信息
	AddCaller bool
	// 日志前缀
	Prefix string
	// 日志输出文件最大长度，超过改值则截断
	MaxSize   int
	MaxAge    int
	MaxBackup int
	// 日志磁盘刷盘间隔
	Interval      time.Duration
	CallerSkip    int
	Async         bool
	Queue         bool
	QueueSleep    time.Duration
	Core          zapcore.Core
	Debug         bool
	EncoderConfig *zapcore.EncoderConfig
}

// Filename 返回文件的绝对路径
func (c Config) Filename() string {
	return fmt.Sprintf("%s/%s", c.Dir, c.Name)
}

// DefaultConfig 应用使用的默认配置
func DefaultConfig() *Config {
	return &Config{
		Name:          "go_kit_default.json",
		Dir:           goKit.LogDir(),
		Level:         "info",
		MaxSize:       500, // 500M
		MaxAge:        1,   // 1 day
		MaxBackup:     100, // 10 backup
		Interval:      24 * time.Hour,
		CallerSkip:    0,
		AddCaller:     true,
		Async:         true,
		Queue:         false,
		QueueSleep:    100 * time.Millisecond,
		Debug:         true,
		EncoderConfig: DefaultZapConfig(),
		Fields: []zap.Field{
			String("appName", goKit.Name()),
			String("appID", goKit.AppID()),
			String("instanceID", goKit.AppInstance()),
		},
	}
}

// Build 构造 Logger
func (c Config) Build() *Logger {
	if c.EncoderConfig == nil {
		c.EncoderConfig = DefaultZapConfig()
	}
	if c.Debug {
		c.EncoderConfig.EncodeLevel = DebugEncodeLevel
	}
	logger := newLogger(&c)

	return logger
}
