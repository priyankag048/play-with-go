package main

import (
	"fmt"
	"sync"
)

var messages = [][]string{
	{
		"The world itself's",
		"just one big hoax.",
		"Spamming each other with our",
		"running commentary of bullshit,",
	},
	{
		"but with our things, our property, our money.",
		"I'm not saying anything new.",
		"We all know why we do this,",
		"not because Hunger Games",
		"books make us happy,",
	},
	{
		"masquerading as insight, our social media",
		"faking as intimacy.",
		"Or is it that we voted for this?",
		"Not with our rigged elections,",
	},
	{
		"but because we wanna be sedated.",
		"Because it's painful not to pretend,",
		"because we're cowards.",
		"- Elliot Alderson",
		"Mr. Robot",
	},
}

const producerCount = 4
const consumerCount = 3

func main() {
	c := make(chan string)
	wp := sync.WaitGroup{}
	wc := sync.WaitGroup{}
	for i := 0; i < producerCount; i++ {
		wp.Add(1)
		go produce(c, i, &wp)
	}
	for i := 0; i < consumerCount; i++ {
		wc.Add(1)
		go consumer(c, i, &wc)
	}
	wp.Wait()
	close(c)
	wc.Wait()
}

func produce(ch chan string, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	// fmt.Println(len(ch))
	for _, msg := range messages[index] {
		ch <- msg
		fmt.Printf("Message from Producer %v is %v\n", index, msg)
	}
}

func consumer(ch chan string, index int, wc *sync.WaitGroup) {
	defer wc.Done()
	for msg := range ch {
		fmt.Printf("Consumed message %v by Consumer %v\n", msg, index)
	}
}
