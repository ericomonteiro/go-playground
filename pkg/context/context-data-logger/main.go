package main

import (
	"context"
	"errors"
	context_logger "go-playground/pkg/context/context-logger"
	"go.uber.org/zap"
	"time"
)

func main() {
	url := "https://example.com"
	logger, _ := zap.NewProduction()
	logger.Info("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger.Info("Failed to fetch URL: %s", zap.String("url", url))

	ctx := context_logger.LoggerToContext(context.Background(), logger)
	requestData := &context_logger.RequestData{
		UserID:    "user-01",
		RequestID: "req-01",
		SiteID:    "MLB",
	}

	ctx = context_logger.NewContextWithData(ctx, requestData)

	context_logger.Info(ctx, "demo log message")

	anotherFunc(ctx)
}

func anotherFunc(ctx context.Context) {
	//doSomething with error
	err := errors.New("unexpected error")

	context_logger.Error(ctx, "demo error message", err)
}
