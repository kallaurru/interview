package algo_206

func ReverseListRecursive(in []int, r, l int) []int {
	if len(in) == 0 {
		return in
	}
	if r < l {
		in[r], in[l] = in[l], in[r]
		r++
		l = len(in) - r - 1
		return ReverseListRecursive(in, r, l)
	}
	return in
}

func ReverseListIteration(in []int) []int {
	if len(in) == 0 {
		return in
	}
	r := 0
	l := len(in) - 1
	for r < l {
		in[r], in[l] = in[l], in[r]
		r++
		l = len(in) - r - 1
	}

	return in
}
