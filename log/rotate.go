package log

import (
	"io"

	"github.com/yidianyipie/go-kit/log/rotate"
)

// newRotate 一个滚动记录日志的文件LogWriter
func newRotate(config *Config) io.Writer {
	if config.Dir == "" {
		config.Dir = "/data/dataLogs/"
	}
	rotateLog := rotate.NewLogger()
	rotateLog.Filename = config.Filename()
	rotateLog.MaxSize = config.MaxSize // MB
	rotateLog.MaxAge = config.MaxAge   // days
	rotateLog.MaxBackups = config.MaxBackup
	rotateLog.Interval = config.Interval
	rotateLog.LocalTime = true
	rotateLog.Compress = false
	return rotateLog
}
