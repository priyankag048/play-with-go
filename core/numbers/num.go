package main

import "fmt"

type numSlice []int

func checkEvenOdd(nums []int) {
	fmt.Println("---- checkEvenOdd ----")
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}

func (n numSlice) reverse() {
	fmt.Println("---- reverse number ----")
	for i := 0; i < len(n)/2; i++ {
		n[i], n[len(n)-i-1] = n[len(n)-i-1], n[i]
	}
}

func perfectSquare(n int) {
	i := 1
	sum := 0
	result := false
	for i < n {
		sum = sum + i
		if sum == n {
			result = true
		}
		i = i + 2
	}
	if result {
		fmt.Println(n, " is a perfect square")
	} else {
		fmt.Println(n, " is not a perfect square")
	}
}
