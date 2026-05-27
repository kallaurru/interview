package algo_350

func IntersectionTwoArraysAlgo350(l []int, r []int) []int {
	if len(l) == 0 && len(r) == 0 {
		return []int{}
	}
	if len(l) == 0 {
		return r
	}
	if len(r) == 0 {
		return l
	}
	// делаем l всегда с бОльшей длиной, а r всегда с меньшей
	if len(l) < len(r) {
		l, r = r, l
	}
	stor := make(map[int]int, len(r))
	result := make([]int, 0, len(r))
	for idx := 0; idx < len(r); idx++ {
		_, ok := stor[r[idx]]
		if ok {
			stor[r[idx]] += 1
		} else {
			stor[r[idx]] = 1
		}
	}

	for i := 0; i < len(l); i++ {
		count, ok := stor[l[i]]
		if !ok || count == 0 {
			continue
		}
		stor[l[i]] -= 1
		result = append(result, l[i])
	}

	return result
}
