package algo_26

func RemoveDuplicateSortedArray26(in []int) int {
	out := make([]int, 0, len(in)/2)
	if len(in) == 1 {
		return 1
	}
	out = append(out, in[0])
	for i := 1; i < len(in); i++ {
		if in[i] == out[len(out)-1] {
			continue
		}
		out = append(out, in[i])
	}

	return len(out)
}
