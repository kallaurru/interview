package algo_3

func LongestSubstring(in string) int {
	l := 0
	out := 0
	stor := make(map[uint8]int, 128)

	for r := 0; r < len(in); r++ {
		stor[in[r]] += 1
		count := stor[in[r]]
		if count < 2 {
			continue
		}

		// обновляем out
		if r-l > out {
			out = r - l
		}
		for l <= r {
			count = stor[in[l]]
			if count > 1 {
				stor[in[l]] -= 1
				break
			}
			l++
		}
		l++
	}
	return out
}
