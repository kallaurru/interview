package algo_1

func TwoSum(in []int, target int) []int {
	if len(in) < 2 {
		return []int{}
	}
	ws := 2
	r, l := in[0], in[1]
	sum := r + l
	if sum == target {
		return []int{0, 1}
	}
	for i := 2; i < len(in); i++ {
		sum += in[i] - in[i-ws]
		if sum == target {
			return []int{}
		}
	}
	return nil
}
