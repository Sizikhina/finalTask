package find_missing_element

import "testing"

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
				[]int{2, 3, 1, 5},
			},
			4,
		},
		{
			args{
				[]int{1, 2, 3, 4, 5, 7, 8, 9, 10},
			},
			6,
		},
		{
			args{
				[]int{2},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Solution(tt.args.A); got != tt.want {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
