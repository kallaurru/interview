package algo_704

func BinSearchAlgo704(in []int, target int) int {
	l := 0
	h := len(in) - 1
	for l <= h {
		m := l + ((h - l) / 2)
		val := in[m]
		if val == target {
			return m
		}
		if val < target {
			l = m + 1
			continue
		}
		if val > target {
			h = m - 1
			continue
		}
	}
	return -1
}
