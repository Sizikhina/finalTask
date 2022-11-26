package entry_into_the_array_test

import (
	"testing"
	"zlata/entry_into_the_array"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{
				[]int{9, 3, 9, 3, 9, 7, 9},
			},
			7,
		},
		{
			args{
				[]int{1, 4, 5, 5, 1, 1, 8, 1, 8},
			},
			4,
		},
		{
			args{
				[]int{10, 2, 9, 2, 2, 2, 9, 5, 5},
			},
			10,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := entry_into_the_array.Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
