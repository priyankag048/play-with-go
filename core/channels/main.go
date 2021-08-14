package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
		"http://google.com",
	}

	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for {
		fmt.Println(<-c)
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}
	fmt.Println(link, "is up")
	c <- link
}
