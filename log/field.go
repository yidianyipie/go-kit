package log

import (
	"fmt"
	"strings"
	"time"

	"go.uber.org/zap"
)

// FieldAid 应用唯一标识符
func FieldAid(value string) Field {
	return String("aid", value)
}

// FieldMod 模块
func FieldMod(value string) Field {
	value = strings.Replace(value, " ", ".", -1)
	return String("mod", value)
}

// FieldAddr 地址。以mysql为例，"dsn = "root:juno@tcp(127.0.0.1:3306)/juno?charset=utf8"，addr为 "127.0.0.1:3306"
func FieldAddr(value string) Field {
	return String("addr", value)
}

// FieldAddrAny 任意地址信息
func FieldAddrAny(value interface{}) Field {
	return Any("addr", value)
}

// FieldName 名字
func FieldName(value string) Field {
	return String("name", value)
}

// FieldType 类型
func FieldType(value string) Field {
	return String("type", value)
}

// FieldCode code
func FieldCode(value int32) Field {
	return Int32("code", value)
}

// FieldCost 耗时时间
func FieldCost(value time.Duration) Field {
	return String("cost", fmt.Sprintf("%.3f", float64(value.Round(time.Microsecond))/float64(time.Millisecond)))
}

// FieldKey str的key值
func FieldKey(value string) Field {
	return String("key", value)
}

// FieldKeyAny 任意类型的key值
func FieldKeyAny(value interface{}) Field {
	return Any("key", value)
}

// FieldValue str的value值
func FieldValue(value string) Field {
	return String("value", value)
}

// FieldValueAny 任意类型的value值
func FieldValueAny(value interface{}) Field {
	return Any("value", value)
}

// FieldErrKind str的errKind值
func FieldErrKind(value string) Field {
	return String("errKind", value)
}

// FieldErr error类型的zapField
func FieldErr(err error) Field {
	return zap.Error(err)
}

// FieldStringErr str类型的error Field
func FieldStringErr(err string) Field {
	return String("err", err)
}

// FieldExtMessage 任意类型的扩展信息，key为ext
func FieldExtMessage(values ...interface{}) Field {
	return zap.Any("ext", values)
}

// FieldStack stack Field
func FieldStack(value []byte) Field {
	return ByteString("stack", value)
}

// FieldMethod 方法Field
func FieldMethod(value string) Field {
	return String("meth", value)
}

// FieldEvent 事件Field
func FieldEvent(value string) Field {
	return String("event", value)
}

// FieldHost 主机Field
func FieldHost(value string) Field {
	return String("host", value)
}

// FieldReqAID 请求的应用ID Field
func FieldReqAID(value string) Field {
	return String("reqAid", value)
}

// FieldIP IP Field
func FieldIP(value string) Field {
	return String("ip", value)
}

// FieldReqHost 请求的主机 Field
func FieldReqHost(value string) Field {
	return String("reqHost", value)
}

// FieldColor color Field
func FieldColor(value string) Field {
	return String("color", value)
}

// FieldStdMethod 标准方法Field
func FieldStdMethod(value string) Field {
	return String("stdMeth", value)
}
