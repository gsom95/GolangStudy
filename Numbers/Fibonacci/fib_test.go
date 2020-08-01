package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input    int
	expected int
}{
	{0, 0},
	{1, 1},
	{4, 3},
	{11, 89},
}

func TestFibSimple(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %d", test.input), func(t *testing.T) {
			got := fib(test.input)
			if got != test.expected {
				t.Errorf("Got: %d, expected %d\n", got, test.expected)
			}
		})
	}

}

func TestFibClosure(t *testing.T) {
	for _, test := range tests {
		f := fibonacci()
		for i := 0; i < test.input; i++ {
			f()
		}
		got := f()
		if got != test.expected {
			t.Errorf("Got: %d, expected %d\n", got, test.expected)
		}
	}
}
