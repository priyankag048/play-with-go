package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main3() {
	checkName()
	checkNumber()
}

func checkName() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("you name is ", name)
}

func checkNumber() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number")
	numberStr, _ := reader.ReadString('\n')
	println("numberStr ", numberStr)
	number, _ := strconv.Atoi(strings.TrimSpace(numberStr))
	println("number ", number)
	if number > 0 && number <= 10 {
		fmt.Println("The number is in between 1 and 10")
	} else {
		fmt.Println("The number is not in between 1 and 10")
	}
}
