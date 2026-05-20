package algo_88

func MergeArraysAlgo88(first, second []int, m, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for idx, val := range second {
			first[idx] = val
		}
		return
	}
	zeroVal := 0
	l, r, s := m-1, len(first)-1, n-1
	for {
		if s < 0 {
			break
		}
		if l < 0 {
			l = 0
		}
		if first[l] <= second[s] {
			first[r] = second[s]
			r--
			s--
			continue
		}
		if first[l] > second[s] {
			first[r] = first[l]
			first[l] = zeroVal
			r--
			l--
			continue
		}
	}
}
