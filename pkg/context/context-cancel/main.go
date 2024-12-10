package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

type ctxCancelFuncKey struct{}

func createContextWithCancel() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, ctxCancelFuncKey{}, cancel)
	return ctx
}

func cancelContext(ctx context.Context) {
	if cancel, ok := ctx.Value(ctxCancelFuncKey{}).(context.CancelFunc); ok {
		cancel()
	}
}

func main() {
	ctx := createContextWithCancel()

	go randomIntAndStopWithEquals(ctx, "GO 01", 5)
	go randomIntAndStopWithEquals(ctx, "GO 02", 9)

	for {
		time.Sleep(100 * time.Millisecond)
		select {
		case <-ctx.Done():
			fmt.Println("Context is done")
			return
		}
	}
}

func randomIntAndStopWithEquals(ctx context.Context, prefix string, expectedInt int) {
	for {
		time.Sleep(100 * time.Millisecond)

		select {
		case <-ctx.Done():
			return
		default:
			randomInt := rand.IntN(10)
			fmt.Printf("%s: %d\n", prefix, randomInt)
			if randomInt == expectedInt {
				cancelContext(ctx)
			}
		}
	}
}
