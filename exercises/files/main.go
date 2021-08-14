package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("./data.txt"); os.IsNotExist(err) {
		fmt.Printf("File does not exist\n")
	}
}
