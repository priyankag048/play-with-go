package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting((eb))
	// printGreeting(sp)
}

func (englishBot) getGreeting() string {
	// very custom logic for generation english greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

// func printGreeting(sp englishBot) {
// 	fmt.Println(sp.getGreeting())
// }
