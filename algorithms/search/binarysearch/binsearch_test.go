package binarysearch

import "testing"

type args struct {
	arr  []int
	find int
}
type testCase struct {
	name string
	args args
	want int
}

var tests = []testCase{
	{
		name: "Empty array",
		args: args{
			arr:  []int{},
			find: 0,
		},
		want: -1,
	},
	{
		name: "Successful search 1",
		args: args{
			arr:  []int{1, 2, 3, 5, 7, 9, 13, 22, 48, 56, 90},
			find: 3,
		},
		want: 2,
	},
	{
		name: "Successful search 2",
		args: args{
			arr:  []int{1, 2, 3, 5, 7, 9, 13, 22, 48, 56, 90},
			find: 22,
		},
		want: 7,
	},
	{
		name: "Successful search 3 (same elements)",
		args: args{
			arr:  []int{1, 1, 3, 5, 5, 9, 13, 22, 48, 56, 90},
			find: 5,
		},
		want: 3,
	},
	{
		name: "Unsuccessful search 1",
		args: args{
			arr:  []int{1, 2, 3, 5, 7, 9, 13, 22, 48, 56, 90},
			find: 23,
		},
		want: -1,
	},
	{
		name: "Unsuccessful search 2",
		args: args{
			arr:  []int{1, 2, 3, 5, 7, 9, 13, 22, 48, 56, 90},
			find: 4,
		},
		want: -1,
	},
}

func TestBinarySearchRecursive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearchRecursive(tt.args.arr, tt.args.find, 0, len(tt.args.arr)); got != tt.want {
				t.Errorf("BinarySearchRecursive() = %v, want %v", got, tt.want)
			}
		})
	}
}
