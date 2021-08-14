package main

import "fmt"

type Data []int

func main() {
	data := Data{1, 2, 3, 4}
	fmt.Println(data.add())
}

func (d Data) add() int {
	result := 0
	for _, num := range d {
		result += num
	}
	return result
}
