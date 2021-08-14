package main

import "fmt"

func fibonacci() func(int) int {
	fib := []int{}
	return func(n int) int {
		if n == 0 || n == 1 {
			fib = append(fib, n)
			return n
		} else if n == 2 {
			fib = append(fib, 1)
			return 1
		} else {
			data := fib[n-1] + fib[n-2]
			fib = append(fib, data)
			return data
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
