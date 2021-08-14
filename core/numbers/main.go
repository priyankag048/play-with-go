package main

import "fmt"

func main() {
	nums := numSlice{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	checkEvenOdd(nums)
	nums.reverse()
	fmt.Println(nums)
	perfectSquare(4)
}
