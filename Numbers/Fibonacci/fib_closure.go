package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
// solution for https://tour.golang.org/moretypes/26
func fibonacci() func() int {
	current := 0
	next := 0
	return func() int {
		if next == 0 {
			next = 1
		} else {
			current, next = next, current+next
		}
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
