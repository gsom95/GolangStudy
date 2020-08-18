package primenumbers

import (
	"fmt"
	"testing"
)

var tests = []struct {
	input    int
	expected []int
}{
	{0, []int{}},
	{2, []int{2}},
	{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
	{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
}

func TestFindPrimeNumbers(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %d", test.input), func(t *testing.T) {
			got := FindPrimeNumbers(test.input)
			if got == nil {
				t.Error("Result is nil")
			}

			if len(got) != len(test.expected) {
				t.Error("Different length in result and expected output")
			}

			for i := range got {
				if got[i] != test.expected[i] {
					t.Errorf("Expected: %d, but got: %d\n", test.expected[i], got[i])
				}
			}
		})

	}
}

func TestSieve(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %d", test.input), func(t *testing.T) {
			got := Sieve(test.input)
			if got == nil {
				t.Fatal("Result is nil")
			}

			if len(got) != len(test.expected) {
				t.Fatal("Different length in result and expected output")
			}

			for i := range got {
				if got[i] != test.expected[i] {
					t.Fatalf("Expected: %d, but got: %d\n", test.expected[i], got[i])
				}
			}
		})

	}
}
