package gocontext

import (
	"context"
	"log"
)

func subFuncPrintCtxValue(ctx context.Context) {
	log.Println(ctx.Value("a"))
	log.Println(ctx.Value("b"))
}
