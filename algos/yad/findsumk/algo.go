package findsumk

func SumKInArr(in []int, k int) (int, int) {
	prev := 0
	subAmounts := make(map[int]int, len(in)/2)

	for l := 0; l < len(in); l++ {
		result := (prev + in[l]) % k
		idx, ok := subAmounts[result]
		if ok {
			return idx + 1, l
		}
		subAmounts[result] = l
		prev = result
	}

	return 0, 0
}
