package main

import (
	"context"
	"fmt"
)

type DataContext struct {
	UserID    string
	RequestID string
	SiteID    string
}

type contextDataKey struct{}

func NewContextWithData(ctx context.Context, data *DataContext) context.Context {
	return context.WithValue(ctx, contextDataKey{}, data)
}

func GetDataFromContext(ctx context.Context) *DataContext {
	if data, ok := ctx.Value(contextDataKey{}).(*DataContext); ok {
		return data
	}
	return nil
}

func main() {
	// Create a empty context
	ctx := context.Background()

	// Simulate context data
	data := &DataContext{
		UserID:    "user-123",
		RequestID: "req-123",
		SiteID:    "MLB",
	}

	// Add context data to context
	ctx = NewContextWithData(ctx, data)

	// Do something with context data
	doSomethingWithCtxData(ctx)
}

func doSomethingWithCtxData(ctx context.Context) {
	data := GetDataFromContext(ctx)
	if data != nil {
		fmt.Println("UserID: ", data.UserID)
		fmt.Println("RequestID: ", data.RequestID)
		fmt.Println("SiteID: ", data.SiteID)
	}
}
