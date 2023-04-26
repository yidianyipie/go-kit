package log

import (
	"context"
	"go.uber.org/multierr"
	"go.uber.org/zap/zapcore"
	"log"

	"go.uber.org/zap"
)

// L 从context中获取Logger。如果不存在就返回defaultLogger
func L(ctx context.Context) *Logger {
	return FromContext(ctx)
}

// GoKit 返回框架使用Logger
func GoKit() *Logger {
	return goKitLogger
}

// SetGoKit 设置框架使用Logger
func SetGoKit(logger *Logger) {
	goKitLogger = logger
}

// Default 返回默认日志记录器
func Default() *Logger {
	return defaultLogger
}

// SetDefault 设置默认Logger，同时会替代zap默认日志和goKit的标准Logger
func SetDefault(logger *Logger) {
	defaultLogger = logger
	stdLogger = defaultLogger.WithOptions(zap.AddCallerSkip(1))
	zap.ReplaceGlobals(defaultLogger)
	zap.RedirectStdLog(defaultLogger)
}

// Debug 记录一条Debug级别的消息。该消息包括日志站点上传递的所有字段，以及日志记录器上累积的所有字段。
func Debug(msg string, fields ...Field) {
	stdLogger.Debug(msg, fields...)
}

// Info 记录一条Info级别的消息。该消息包括日志站点上传递的所有字段，以及日志记录器上累积的所有字段。
func Info(msg string, fields ...Field) {
	stdLogger.Info(msg, fields...)
}

// Warn 记录一条Warn级别的消息。该消息包括日志站点上传递的所有字段，以及日志记录器上累积的所有字段。
func Warn(msg string, fields ...Field) {
	stdLogger.Warn(msg, fields...)
}

// Error 记录一条Error级别的消息。该消息包括日志站点上传递的所有字段，以及日志记录器上累积的所有字段。
func Error(msg string, fields ...Field) {
	stdLogger.Error(msg, fields...)
}

// DPanic 记录一条DPanic级别的消息。该消息包括日志站点上传递的所有字段，以及日志记录器上累积的所有字段。
func DPanic(msg string, fields ...Field) {
	stdLogger.DPanic(msg, fields...)
}

// Panic 记录一条Panic级别的消息。该消息包括日志站点上传递的所有字段，以及日志记录器上累积的所有字段。
func Panic(msg string, fields ...Field) {
	stdLogger.Panic(msg, fields...)
}

// Fatal 记录一条Fatal级别的消息。该消息包括日志站点上传递的所有字段，以及日志记录器上累积的所有字段。
func Fatal(msg string, fields ...Field) {
	stdLogger.Fatal(msg, fields...)
}

// DebugW 记录一条Debug级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，以及日志记录器上累积的所有字段。
func DebugW(msg string, keysAndValues ...interface{}) {
	fields := sweetenFields(keysAndValues)
	println("fields: ", len(fields))
	stdLogger.Debug(msg, fields...)
}

// InfoW 记录一条Info级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，以及日志记录器上累积的所有字段。
func InfoW(msg string, keysAndValues ...interface{}) {
	fields := sweetenFields(keysAndValues)
	stdLogger.Info(msg, fields...)
}

// WarnW 记录一条Warn级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，以及日志记录器上累积的所有字段。
func WarnW(msg string, keysAndValues ...interface{}) {
	fields := sweetenFields(keysAndValues)
	stdLogger.Warn(msg, fields...)
}

// ErrorW 记录一条Error级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，以及日志记录器上累积的所有字段。
func ErrorW(msg string, keysAndValues ...interface{}) {
	fields := sweetenFields(keysAndValues)
	stdLogger.Error(msg, fields...)
}

// DPanicW 记录一条DPanic级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，以及日志记录器上累积的所有字段。
func DPanicW(msg string, keysAndValues ...interface{}) {
	fields := sweetenFields(keysAndValues)
	stdLogger.DPanic(msg, fields...)
}

// PanicW 记录一条Panic级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，以及日志记录器上累积的所有字段。
func PanicW(msg string, keysAndValues ...interface{}) {
	fields := sweetenFields(keysAndValues)
	stdLogger.Panic(msg, fields...)
}

// FatalW 记录一条Fatal级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，以及日志记录器上累积的所有字段。
func FatalW(msg string, keysAndValues ...interface{}) {
	fields := sweetenFields(keysAndValues)
	stdLogger.Fatal(msg, fields...)
}

// DebugC 记录一条带上下文的Debug级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，trace_id追踪字段，以及日志记录器上累积的所有字段。
func DebugC(ctx context.Context, msg string, keysAndValues ...interface{}) {
	var fields = make([]Field, 0, len(keysAndValues)+1)
	fields = append(fields, Any(TraceIDField, ctx.Value(TraceIDField)))
	fields = append(fields, sweetenFields(keysAndValues)...)
	stdLogger.Debug(msg, fields...)
}

// InfoC 记录一条带上下文的Info级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，trace_id追踪字段，以及日志记录器上累积的所有字段。
func InfoC(ctx context.Context, msg string, keysAndValues ...interface{}) {
	var fields = make([]Field, 0, len(keysAndValues)+1)
	fields = append(fields, Any(TraceIDField, ctx.Value(TraceIDField)))
	fields = append(fields, sweetenFields(keysAndValues)...)
	stdLogger.Info(msg, fields...)
}

// WarnC 记录一条带上下文的Warn级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，trace_id追踪字段，以及日志记录器上累积的所有字段。
func WarnC(ctx context.Context, msg string, keysAndValues ...interface{}) {
	var fields = make([]Field, 0, len(keysAndValues)+1)
	fields = append(fields, Any(TraceIDField, ctx.Value(TraceIDField)))
	fields = append(fields, sweetenFields(keysAndValues)...)
	stdLogger.Warn(msg, fields...)
}

// ErrorC 记录一条带上下文的Error级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，trace_id追踪字段，以及日志记录器上累积的所有字段。
func ErrorC(ctx context.Context, msg string, keysAndValues ...interface{}) {
	var fields = make([]Field, 0, len(keysAndValues)+1)
	fields = append(fields, Any(TraceIDField, ctx.Value(TraceIDField)))
	fields = append(fields, sweetenFields(keysAndValues)...)
	stdLogger.Error(msg, fields...)
}

// DPanicC 记录一条带上下文的DPanic级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，trace_id追踪字段，以及日志记录器上累积的所有字段。
func DPanicC(ctx context.Context, msg string, keysAndValues ...interface{}) {
	var fields = make([]Field, 0, len(keysAndValues)+1)
	fields = append(fields, Any(TraceIDField, ctx.Value(TraceIDField)))
	fields = append(fields, sweetenFields(keysAndValues)...)
	stdLogger.DPanic(msg, fields...)
}

// PanicC 记录一条带上下文的Panic级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，trace_id追踪字段，以及日志记录器上累积的所有字段。
func PanicC(ctx context.Context, msg string, keysAndValues ...interface{}) {
	var fields = make([]Field, 0, len(keysAndValues)+1)
	fields = append(fields, Any(TraceIDField, ctx.Value(TraceIDField)))
	fields = append(fields, sweetenFields(keysAndValues)...)
	stdLogger.Panic(msg, fields...)
}

// FatalC 记录一条带上下文的Fatal级别的消息。根据传递的kv信息构造zap field。该消息包括这些字段，trace_id追踪字段，以及日志记录器上累积的所有字段。
func FatalC(ctx context.Context, msg string, keysAndValues ...interface{}) {
	var fields = make([]Field, 0, len(keysAndValues)+1)
	fields = append(fields, Any(TraceIDField, ctx.Value(TraceIDField)))
	fields = append(fields, sweetenFields(keysAndValues)...)
	stdLogger.Fatal(msg, fields...)
}

// With 创建一个子logger并为其添加结构化上下文。添加到子元素的字段不会影响父元素，反之亦然。
func With(fields ...Field) *Logger {
	return stdLogger.With(fields...)
}

// WithOptions 克隆当前的Logger，应用提供的选项，并返回结果Logger并发使用是安全的。
func WithOptions(opts ...Option) *Logger {
	return stdLogger.WithOptions(opts...)
}

// Named 将一个新的路径段添加到logger的名称中段由句点连接。默认情况下，logger是未命名的。
func Named(s string) *Logger {
	return stdLogger.Named(s)
}

// sweetenFields 根据用户任意输入构造zapcore.Field，并返回
func sweetenFields(args []interface{}) []Field {
	if len(args) == 0 {
		return nil
	}
	var (
		fields    = make([]Field, 0, len(args))
		invalid   invalidPairs
		seenError bool
	)
	for i := 0; i < len(args); {
		if f, ok := args[i].(Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		if err, ok := args[i].(error); ok {
			if !seenError {
				seenError = true
				fields = append(fields, zap.Error(err))
			} else {
				stdLogger.Error(_multipleErrMsg, zap.Error(err))
			}
			i++
			continue
		}

		if i == len(args)-1 {
			stdLogger.Error(_oddNumberErrMsg, Any("ignored", args[i]))
			break
		}

		//使用这个值和下一个值，将它们视为键值对。如果键不是字符串，则将这个对添加到无效对的切片中。
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, Any(keyStr, val))
		}

		i += 2
	}

	if len(invalid) > 0 {
		stdLogger.Error(_nonStringKeyErrMsg, Array("invalid", invalid))
	}

	return fields
}

// SyncDefaultLogger 同步默认日志记录器的缓存日志
func SyncDefaultLogger() {
	if err := defaultLogger.Sync(); err != nil {
		log.Printf("default log synchronization failed. error: %v", err)
	}
}

// SyncStandardLogger 同步标准日志记录器的缓存日志
func SyncStandardLogger() {
	if err := stdLogger.Sync(); err != nil {
		log.Printf("standard log synchronization failed. error: %v", err)
	}
}

const (
	// _oddNumberErrMsg 忽略没有值的键。
	_oddNumberErrMsg = "Ignored key without a value."
	// _nonStringKeyErrMsg 忽略非字符串键的键值对。
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
	// _multipleErrMsg 没有键的多重错误。
	_multipleErrMsg = "Multiple errors without a key."
)

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	Any("key", p.key).AddTo(enc)
	Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}
