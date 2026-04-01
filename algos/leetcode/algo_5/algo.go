package algo_5

func MaxPalindromeAlgo5(in string) string {
	minVal := func(val1, val2 int) int {
		if val1 <= val2 {
			return val1
		}
		return val2
	}
	var (
		center       = 1
		radius       = 0
		isOdd        = true
		l, r, radius = 0, 0, 0
	)

	// записываем радиусы для каждого центра i
	d := make([]int, len(in), len(in))
	for i := 1; i < len(in); i++ {
		if i < r {
			d[i] = minVal(r-i+1, d[l+r-i])
		}
		// пока левая граница не вываливается за 0
		// правая граница не выходит за макс индекс
		// и символы равны между собой на одинаковом расстояния от центра
		for i-d[i] >= 0 && i+d[i] < len(in) && in[i-d[i]] == in[i+d[i]] {
			d[i]++
		}
	}

	return ""
}
