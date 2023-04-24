package errcode

import (
	"fmt"
	"strings"
)

type Error struct {
	code    uint32
	msg     string
	details []string
}

var codeMap = map[uint32]string{}

func New(code uint32, msg string) *Error {
	if packageCode == 0 {
		var packageNames []string
		for p := range PackageNum {
			packageNames = append(packageNames, fmt.Sprintf("[%s]", p))
		}
		panic(fmt.Sprintf("请先设置项目package: %s", strings.Join(packageNames, " ")))
	}
	errCode := systemErrorCode + packageCode + code
	if _, exist := codeMap[errCode]; exist {
		panic(fmt.Sprintf("错误码 %d 已存在，请更换一个", code))
	}
	codeMap[errCode] = msg
	return &Error{
		code: errCode,
		msg:  msg,
	}
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("错误码：%d，错误信息：%s", e.Code(), e.Msg())
}

func (e *Error) Code() uint32 {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) MsgF(args ...interface{}) string {
	if len(args) == 0 {
		return e.Msg()
	}
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newErr := *e
	newErr.details = []string{}
	newErr.details = append(newErr.details, details...)
	return &newErr
}
