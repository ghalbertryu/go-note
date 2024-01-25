package gocontext

import (
	"context"
	"testing"
)

func TestValue(t *testing.T) {
	bgCtx := context.Background()
	subFuncPrintCtxValue(bgCtx)

	valueCtx := context.WithValue(bgCtx, "a", "1")
	subFuncPrintCtxValue(valueCtx)

	value2Ctx := context.WithValue(valueCtx, "b", "2")
	subFuncPrintCtxValue(value2Ctx)
}
