package main

import (
	"context"
	"log"
	"runtime"
	"time"

	"github.com/rakyll/gcptrace"
)

var t *trace.Client

func main() {
	ctx, err := trace.New(context.Background(), "jbd-gce")
	if err != nil {
		log.Fatal(err)
	}
	f(ctx)
}

func f(ctx context.Context) {
	ctx = trace.WithSpan(ctx, "")
	defer trace.Finish(ctx)

	go a1(ctx)
	a2(ctx)
	a3(ctx)
}

func a1(ctx context.Context) {
	ctx = trace.WithSpan(ctx, "")
	defer trace.Finish(ctx)

	trace.Logf(ctx, "this is a format string, num goroutines: %v", runtime.NumGoroutine())
	time.Sleep(100 * time.Millisecond)
}

func a2(ctx context.Context) {
	ctx = trace.WithSpan(ctx, "a2")
	defer trace.Finish(ctx)

	time.Sleep(200 * time.Millisecond)
}

func a3(ctx context.Context) {
	ctx = trace.WithSpan(ctx, "a3")
	defer trace.Finish(ctx)

	time.Sleep(300 * time.Millisecond)
}
