package main

import (
	"fmt"
	"strconv"
)

func fib(n int) int {
	num1, num2 := 0, 1
	if n == 0 {
		return 0
	}

	for i := 1; i < n; i++ {
		num1, num2 = num2, num1+num2
	}

	return num2
}

func main() {
	fmt.Print("Enter index: ")
	var input string
	fmt.Scanln(&input)

	i, err := strconv.Atoi(input)
	if err == nil {
		fmt.Println(fib(i))
	} else {
		fmt.Printf("Error during conversion from str to int: %v\n", err)
	}
}
