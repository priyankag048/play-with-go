package main

import "fmt"

func main2() {
	s := "Hello world\nI am Pri"
	for i := 0; i < len(s); i++ {
		fmt.Printf("key:%d, value:%c === %d\n", i, s[i], s[i])
	}

}
