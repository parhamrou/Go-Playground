package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// naive approach
	websites := []string {
		"http://go.dev",
		"http://stackoverflow.com",
		"http://google.com",
	}

	// Creating channel
	c := make(chan string)

	for _, website := range websites {
		go checkWebsite(website, c)
	}

	// We want to request to these websites infinite times. We use this syntax, which is equivalant to the piece of code that
	// is infront of that

	for l := range c {				
		go func(link string) {
			time.Sleep(time.Second * 2)
			checkWebsite(link, c)
		}(l)		
	}								
}

func checkWebsite(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("There is an issue with", link)
		c <- link
		return
	}
	fmt.Println(link, "looks fine!")
	c <- link
}