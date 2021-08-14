package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f := os.Args[1]
	file, err := os.Open(f)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
