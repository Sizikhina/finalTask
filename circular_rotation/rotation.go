package circular_rotation

func Solution(A []int, K int) []int {
	if len(A) == 0 {
		return A
	}
	l := len(A) - K%len(A)
	return append(A[l:], A[:l]...)
}
