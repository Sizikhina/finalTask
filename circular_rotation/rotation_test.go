package circular_rotation_test

import (
	"reflect"
	"testing"
	"zlata/circular_rotation"
)

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		K int
	}
	tests := []struct {
		args args
		want []int
	}{
		{
			args{
				[]int{3, 8, 9, 7, 6},
				3,
			},
			[]int{9, 7, 6, 3, 8},
		},
		{
			args{
				[]int{1, 2, 3, 4, 5, 6},
				6,
			},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			args{
				[]int{},
				3,
			},
			[]int{},
		},
		{
			args{
				[]int{1},
				6,
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := circular_rotation.Solution(tt.args.A, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
