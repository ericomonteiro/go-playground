package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	response := make(map[string]string)
	ctx := req.Context()
	ctx, cancel := context.WithCancel(ctx)

	// to be uncancelable
	//ctx = context.WithoutCancel(ctx)

	defer cancel()

	ch := make(chan string, 1)

	go doSomething(ctx, ch)

	select {
	case <-ctx.Done():
		fmt.Println("Context cancelled: ", ctx.Err())
	case result := <-ch:
		response["msg"] = result
		responseJson, _ := json.Marshal(response)
		_, _ = w.Write(responseJson)
	}
}

func main() {
	http.HandleFunc("/", hello)
	_ = http.ListenAndServe(":8080", nil)
}

func doSomething(ctx context.Context, ch chan string) {
	maxExecutions := 100
	count := 0

	for {
		select {
		case <-ctx.Done():
			return
		default:
			count++
			time.Sleep(time.Millisecond * 100)
			fmt.Printf("Execution %d\n", count)
			if count >= maxExecutions {
				ch <- "Hello, World!"
				return
			}
		}
	}
}
