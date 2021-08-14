package main

import "fmt"

var messages = []string{
	"The world itself's",
	"just one big hoax.",
	"Spamming each other with our",
	"running commentary of bullshit,",
	"masquerading as insight, our social media",
	"faking as intimacy.",
	"Or is it that we voted for this?",
	"Not with our rigged elections,",
	"but with our things, our property, our money.",
	"I'm not saying anything new.",
	"We all know why we do this,",
	"not because Hunger Games",
	"books make us happy,",
	"but because we wanna be sedated.",
	"Because it's painful not to pretend,",
	"because we're cowards.",
	"- Elliot Alderson",
	"Mr. Robot",
}

const consumerCount int = 3

func main() {
	ch := make(chan string)
	done := make(chan bool)
	// go consume(1, ch, done)
	go produce(ch)
	for i := 1; i <= consumerCount; i++ {
		go consume(i, ch, done)
	}

	<-done
}

func produce(c chan string) {
	for _, m := range messages {
		fmt.Println("Produce message", m)
		c <- m
	}
	close(c)
}

func consume(worker int, c chan string, done chan bool) {
	for msg := range c {
		fmt.Printf("Message %v is comsumed by worker %v. \n", msg, worker)
	}
	done <- true
}
