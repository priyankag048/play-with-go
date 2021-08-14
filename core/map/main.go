package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"white": "#ffffff",
		"black": "#000000",
	}
	colors["green"] = "#erfgth"
	delete(colors, "red")
	elem, ok := colors["red"]
	println("The value", elem, "present? ", ok)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
