package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var ctxTimeout context.Context
	var ch chan string
	var cancel context.CancelFunc

	defer func() {
		if cancel != nil {
			cancel()
		}
	}()

	// Expect timeout ----
	fmt.Println("=== Expect timeout ===")
	// Channel used to receive the result from doSomething function
	ch = make(chan string, 1)

	// Create a context with a timeout of 3 seconds
	ctxTimeout, cancel = context.WithTimeout(context.Background(), time.Second*3)

	// Start the doSomething function
	go doSomething(ctxTimeout, "case 01", 4, ch)

	// Wait for one of the following to happen:
	select {
	case <-ctxTimeout.Done():
		fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	}
	fmt.Println("\n")

	// Expect success with timeout ----
	fmt.Println("=== Expect execute with success (cancelable context) ===")
	// Channel used to receive the result from doSomething function
	ch = make(chan string, 1)

	// Create a context with a timeout of 5 seconds but uncancelable
	ctxTimeout, cancel = context.WithTimeout(context.Background(), time.Second*3)
	ctxTimeout = context.WithoutCancel(ctxTimeout)

	// Start the doSomething function
	go doSomething(ctxTimeout, "case 02", 4, ch)

	// Wait for one of the following to happen:
	select {
	case <-ctxTimeout.Done():
		fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	}
	fmt.Println("\n")

	// Expect success with timeout ----
	fmt.Println("=== Expect execute with success ===")
	// Channel used to receive the result from doSomething function
	ch = make(chan string, 1)

	// Create a context with a timeout of 5
	ctxTimeout, cancel = context.WithTimeout(context.Background(), time.Second*5)
	ctxTimeout = context.WithoutCancel(ctxTimeout)

	// Start the doSomething function
	go doSomething(ctxTimeout, "case 03", 4, ch)

	// Wait for one of the following to happen:
	select {
	case <-ctxTimeout.Done():
		fmt.Printf("Context cancelled: %v\n", ctxTimeout.Err())
	case result := <-ch:
		fmt.Printf("Received: %s\n", result)
	}
	fmt.Println("\n")
}

func expectTimeout(ctx context.Context) {

}

func doSomething(ctx context.Context, prefix string, secondsToWait int, ch chan string) {
	duration := time.Duration(secondsToWait) * time.Second
	fmt.Println(prefix + " doSomething Sleeping...")
	time.Sleep(duration)
	fmt.Println(prefix + " doSomething Wake up...")
	ch <- "Did Something"
}
