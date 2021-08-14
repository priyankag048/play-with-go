package main

import "fmt"

func main() {
	p1 := newPage("TestPage", []byte("This is a sample Page."))
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
