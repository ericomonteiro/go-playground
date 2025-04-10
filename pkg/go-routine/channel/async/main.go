package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.amazon.com",
		"https://www.golang.org",
	}

	c := make(chan string)
	for _, link := range links {
		// go keyword creates a new go routine
		go checkLink(link, c)
	}

	//for i := 0; i < len(links); i++ {
	//	// waits for the channel to return a value
	//	fmt.Println(<-c)
	//}

}

func checkLink(link string, c chan string) {
	r, err := http.Get(link)

	if err != nil {
		c <- fmt.Sprintf("%s is down!", link)
		return
	}

	c <- fmt.Sprintf("%s is up! Status: %d", link, r.StatusCode)
}
