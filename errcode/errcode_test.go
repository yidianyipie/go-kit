package errcode_test

import (
	"practice.com/go-kit/errcode"
	"testing"
)

func TestNewError(t *testing.T) {
	errcode.SetPackageCode("enos")
	internal := errcode.New(1, "internal")
	if internal.Code() != 1120001 {
		t.Logf("expected %d, got: %d", 1120001, internal.Code())
	}
	if internal.Msg() != "internal" {
		t.Logf("expected %s, got: %s", "internal", internal.Msg())
	}
}

func TestError(t *testing.T) {
	var err *errcode.Error
	err.Error()
}
