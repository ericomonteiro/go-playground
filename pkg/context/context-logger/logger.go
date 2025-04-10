package context_logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"runtime"
)

type loggerCtxKey struct{}

type requestDataKey struct{}

type RequestData struct {
	UserID    string
	RequestID string
	SiteID    string
}

func NewContextWithData(ctx context.Context, data *RequestData) context.Context {
	return context.WithValue(ctx, requestDataKey{}, data)
}

func LoggerToContext(ctx context.Context, logger *zap.Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey{}, logger)
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	logger, fields := getLoggerAndAppendedFieldsFromContext(ctx, fields...)
	logger.Info(msg, fields...)
}

func Error(ctx context.Context, msg string, err error, fields ...zap.Field) {
	logger, fields := getLoggerAndAppendedFieldsFromContext(ctx, fields...)
	fields = append(fields, zap.Error(err))
	logger.Error(msg, fields...)
}

func getLoggerAndAppendedFieldsFromContext(ctx context.Context, fields ...zap.Field) (*zap.Logger, []zap.Field) {
	_, file, no, ok := runtime.Caller(2)
	if ok {
		fields = append(fields, zap.String("caller", fmt.Sprintf("%s:%d", file, no)))
	}

	logger, ok := ctx.Value(loggerCtxKey{}).(*zap.Logger)
	if !ok {
		return nil, nil
	}

	if requestData, ok := ctx.Value(requestDataKey{}).(*RequestData); ok {
		fields = append(fields, zap.String("UserID", requestData.UserID))
		fields = append(fields, zap.String("RequestID", requestData.RequestID))
		fields = append(fields, zap.String("SiteID", requestData.SiteID))
	}

	return logger, fields
}
