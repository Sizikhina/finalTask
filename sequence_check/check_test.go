package sequence_check

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
			0,
		},
		{
			args{
				[]int{1, 2, 4, 5, 3},
			},
			1,
		},
		{
			args{
				[]int{1},
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
