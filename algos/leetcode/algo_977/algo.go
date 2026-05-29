package algo_977

func SquaresOfSortedArrays977(in []int) []int {
	result := make([]int, len(in))
	l, r, k := 0, len(in)-1, len(in)-1

	for l != r {
		if in[l]*in[l] <= in[r]*in[r] {
			result[k] = in[r] * in[r]
			r--
		} else {
			result[k] = in[l] * in[l]
			l++
		}
		k--
	}
	result[k] = in[l] * in[l]

	return result
}
