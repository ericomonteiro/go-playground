package main

import (
	"fmt"
	"time"
)

/*
 * Complete the 'BurstyRateLimiter' function below.
 *
 * The function accepts following parameters:
 *  1. chan bool requestChan
 *  2. chan int resultChan
 *  3. int batchSize
 *  4. int toAdd



 */

func BurstyRateLimiter(requestChan chan bool, resultChan chan int, batchSize int, toAdd int) {
	currentValue := 0
	for range requestChan {
		for i := 0; i < batchSize; i++ {
			resultChan <- currentValue
			currentValue += toAdd
		}
	}
}

func main() {
	skipBatches := 0
	printBatches := 4
	batchSize := 2
	toAdd := 3

	resultChan := make(chan int)
	requestChan := make(chan bool)
	go BurstyRateLimiter(requestChan, resultChan, batchSize, toAdd)
	for i := 0; i < skipBatches+printBatches; i++ {
		start := time.Now().UnixNano()
		requestChan <- true
		for j := 0; j < batchSize; j++ {
			newResultChan := <-resultChan
			if i < skipBatches {
				continue
			}
			fmt.Println(newResultChan)
		}
		if i >= skipBatches && i != skipBatches+printBatches-1 {
			fmt.Println("-----")
		}
		end := time.Now().UnixNano()
		timeDiff := (end - start) / 1000000
		if timeDiff > 3 {
			fmt.Println("Rate is too high")
			break
		}
	}
}
