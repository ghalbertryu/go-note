package gocontext

import (
	"context"
	"log"
	"time"
)

func runFunc(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			log.Println("done")
			return
		default:
			log.Println("run")
		}
	}
}
