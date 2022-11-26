package sequence_check

import "sort"

func Solution(A []int) int {
	sort.Ints(A)
	for i, value := range A {
		if value != i+A[0] {
			return 0
		}
	}
	return 1
}
