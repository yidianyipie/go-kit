package log

import (
	"os"
	"time"

	"github.com/fatih/color"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Field  = zap.Field
	Level  = zapcore.Level
	Logger = zap.Logger
	Option = zap.Option
)

var (
	// String ...
	String = zap.String
	// Any ...
	Any = zap.Any
	// Int64 ...
	Int64 = zap.Int64
	// Int ...
	Int = zap.Int
	// Int32 ...
	Int32 = zap.Int32
	// Uint ...
	Uint = zap.Uint
	// Duration ...
	Duration = zap.Duration
	// Durationp ...
	Durationp = zap.Durationp
	// Object ...
	Object = zap.Object
	// Namespace ...
	Namespace = zap.Namespace
	// Reflect ...
	Reflect = zap.Reflect
	// Skip ...
	Skip = zap.Skip()
	// ByteString ...
	ByteString = zap.ByteString
	// Array ...
	Array = zap.Array
	// Time ...
	Time = zap.Time
)

const (
	// defaultBufferSize 设置与每个WriterSync关联的缓冲区大小。
	defaultBufferSize = 256 * 1024

	// defaultFlushInterval 默认的刷写间隔
	defaultFlushInterval = 5 * time.Second
)

// defaultLogger 业务使用的默认日志
// stdLogger 标准日志，库函数(Debug,Info...)使用
// goKitLogger 框架使用
var defaultLogger, stdLogger, goKitLogger *Logger

func init() {
	SetDefault(Config{
		Name:  "default",
		Debug: true,
	}.Build())

	SetGoKit(Config{
		Name:  "framework",
		Debug: true,
	}.Build())
}

// newLogger 根据日志配置生成zap日志
func newLogger(config *Config) *zap.Logger {
	zapOptions := make([]zap.Option, 0)
	zapOptions = append(zapOptions, zap.AddStacktrace(zap.ErrorLevel))
	if config.AddCaller {
		zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(config.CallerSkip))
	}
	if len(config.Fields) > 0 {
		zapOptions = append(zapOptions, zap.Fields(config.Fields...))
	}

	var ws zapcore.WriteSyncer
	if config.Debug {
		ws = os.Stdout
	} else {
		ws = zapcore.AddSync(newRotate(config))
	}

	if config.Async {
		ws = &zapcore.BufferedWriteSyncer{
			WS:            zapcore.AddSync(ws),
			FlushInterval: defaultFlushInterval,
			Size:          defaultBufferSize,
		}
	}

	lv := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	if err := lv.UnmarshalText([]byte(config.Level)); err != nil {
		panic(err)
	}

	encoderConfig := *config.EncoderConfig
	core := config.Core
	if core == nil {
		core = zapcore.NewCore(
			func() zapcore.Encoder {
				if config.Debug {
					return zapcore.NewConsoleEncoder(encoderConfig)
				}
				return zapcore.NewJSONEncoder(encoderConfig)
			}(),
			ws,
			lv,
		)
	}

	zapLogger := zap.New(
		core,
		zapOptions...,
	)

	return zapLogger.Named(config.Name)
}

// DefaultZapConfig 返回zap编码配置
func DefaultZapConfig() *zapcore.EncoderConfig {
	return &zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "lv",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// DebugEncodeLevel 调试编码级别，彩色打印
func DebugEncodeLevel(lv zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var colorize = color.RedString
	switch lv {
	case zapcore.DebugLevel:
		colorize = color.BlueString
	case zapcore.InfoLevel:
		colorize = color.GreenString
	case zapcore.WarnLevel:
		colorize = color.YellowString
	case zapcore.ErrorLevel, zap.PanicLevel, zap.DPanicLevel, zap.FatalLevel:
		colorize = color.RedString
	default:
	}
	enc.AppendString(colorize(lv.CapitalString()))
}
