package main

import (
	"fmt"
	"net/http"
	"time"
)

var links = [5]string{
	"https://google.com",
	"https://facebook.com",
	"https://amazon.com",
	"http://medium.com",
	"http://stackoverflow.com",
}

func main() {
	// blocking call
	checkLink("Blocking: ")

	// call using go routines
	go checkLink("Go routines")

	// call from go func
	go func(msg string) {
		checkLink("from go func:" + msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("Done!")
}

func checkLink(from string) {
	for i, link := range links {
		fmt.Println(from, ":", i)
		http.Get(link)
	}
}
