package algo_1

// TwoSumNear Исходя из представленных тест кейсов ищем сумму 2 ух чисел рядом
func TwoSumNear(in []int, target int) []int {
	if len(in) < 2 {
		return []int{}
	}
	ws := 2
	sum := in[0] + in[1]
	if sum == target {
		return []int{0, 1}
	}
	for l := 2; l < len(in); l++ {
		sum += in[l] - in[l-ws]
		if sum == target {
			return []int{l - 1, l}
		}
	}

	return nil
}
