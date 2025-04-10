package main

import (
	"fmt"
	"time"
)

func channelListener(listener chan bool) {
	// Infinite loop to listen to the channel
	for b := range listener {
		fmt.Println("Listener received a message ", b)
	}

	fmt.Println("Listener is done listening")
}

func main() {
	channel := make(chan bool)
	go channelListener(channel)

	channel <- true
	time.Sleep(1 * time.Second)

	channel <- true
	time.Sleep(2 * time.Second)

	channel <- true
	time.Sleep(1 * time.Second)

	close(channel)
	time.Sleep(1 * time.Second)
}
