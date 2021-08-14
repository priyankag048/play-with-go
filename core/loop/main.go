package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	for _, value := range nums {
		fmt.Println("The value is", value)
	}
	c := checkValue(1)
	fmt.Println("The returned value from checkedValue is", c)

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("index", i)
	}
}

func checkValue(x int) float64 {
	if x > 0 {
		return float64(x) / 2
	} else {
		return float64(x)
	}
}
