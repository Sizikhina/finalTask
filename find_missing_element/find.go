package find_missing_element

import "sort"

func Solution(A []int) int {
	sort.Ints(A)
	for i, value := range A {
		if value != i+1 {
			return value - 1
		}
	}
	return 0
}
