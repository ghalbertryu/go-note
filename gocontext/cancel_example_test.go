package gocontext

import (
	"context"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	go runTimeTicker(ctx)

	time.Sleep(3 * time.Second)
	cancelFunc()
	time.Sleep(3 * time.Second)
}

func TestDeadline(t *testing.T) {
	ctx, cancelFunc := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancelFunc()
	go runTimeTicker(ctx)

	time.Sleep(5 * time.Second)
}

func TestTimeout(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFunc()
	go runTimeTicker(ctx)

	time.Sleep(5 * time.Second)
}
