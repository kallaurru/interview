package algo_283

func MoveZero(in []int) []int {
	endGroup := 0
	for i := 0; i < len(in); i++ {
		if in[i] == 0 {
			continue
		}
		in[endGroup], in[i] = in[i], in[endGroup]
		endGroup += 1
	}

	return in
}
