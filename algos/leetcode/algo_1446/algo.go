package algo_1446

func ConsecutiveChar(in string) int {
	maxVal := 1
	counter := 1
	for idx := 1; idx < len(in); idx++ {
		if in[idx-1] == in[idx] {
			counter++
			continue
		}
		if counter > maxVal {
			maxVal = counter
		}
		counter = 1
	}

	return maxVal
}
