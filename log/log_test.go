package log_test

import (
	"context"
	"github.com/yidianyipie/go-kit/log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func Test_log(t *testing.T) {
	stdLog := log.Default()
	stdLog.Debug("debug", log.Any("a", "b"))
	stdLog.Info("info", log.Any("a", "b"))
	stdLog.Warn("warn", log.Any("a", "b"))
	stdLog.Error("error", log.Any("a", "b"))
}

func Test_trace(t *testing.T) {
	kitLog := log.GoKit()
	ctx := log.NewContext(context.TODO(), kitLog, "a:b:c:1")

	stdLog := log.FromContext(ctx)
	stdLog.Debug("debug", log.Any("a", "b"))
	stdLog.Info("info", log.Any("a", "b"))
	stdLog.Warn("warn", log.Any("a", "b"))
	stdLog.Error("error", log.Any("a", "b"))
}

func TestDebugC(t *testing.T) {
	log.DebugW("debug", "key1", "value1")
	log.InfoW("info", "keyInfo", "valueInfo")
	log.WarnW("warn", "keyWarn", "valueWarn", log.String("testFieldKey", "testFieldValue"))
	log.WarnW("warn", "keyWarn", "valueWarn", log.String("testFieldKey", "testFieldValue"), "withoutValue")
	log.ErrorW("error", "keyError", "valueError")
}

func TestLog(t *testing.T) {
	defaultConfig := log.DefaultConfig()
	core, oLog := observer.New(zapcore.InfoLevel)
	defaultConfig.Core = core
	defaultLog := defaultConfig.Build()

	defaultLog.Debug("debug", log.Any("a", "b"))
	defaultLog.Info("info", log.Any("a", "b"), log.FieldCost(time.Second))
	defaultLog.Warn("warn", log.Any("a", "b"))
	defaultLog.Error("error", log.Any("a", "b"))

	assert.Equal(t, 3, len(oLog.All()))
	assert.Equal(t, "info", oLog.All()[0].Message)
	assert.Equal(t, "b", oLog.All()[0].ContextMap()["a"])
	assert.Equal(t, "1000.000", oLog.All()[0].ContextMap()["cost"])
	assert.Equal(t, "1234567890", oLog.All()[0].ContextMap()["appID"])
}
