package algo_268

func MissingNumberAlgo268(in []int) int {
	n := len(in)
	expected := n * (n + 1) / 2
	actual := 0
	for _, val := range in {
		actual += val
	}

	return expected - actual
}
