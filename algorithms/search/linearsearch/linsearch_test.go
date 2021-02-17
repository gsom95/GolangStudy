package linearsearch

import "testing"

func TestLinearSearch(t *testing.T) {
	type args struct {
		arr  []int
		find int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Empty array",
			args: args{
				arr:  []int{},
				find: 1,
			},
			want: -1,
		},
		{
			name: "Successful search",
			args: args{
				arr:  []int{3, 2, 1, 50, 68, 9},
				find: 1,
			},
			want: 2,
		},
		{
			name: "Unsuccessful search",
			args: args{
				arr:  []int{3, 1, 5, 8, 10, 20, 100},
				find: 109,
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LinearSearch(tt.args.arr, tt.args.find); got != tt.want {
				t.Errorf("LinearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
