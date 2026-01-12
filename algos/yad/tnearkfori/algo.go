package tnearkfori

func NearKElements(data []int, index, k int) []int {
	out := make([]int, 0, k)
	// крайние случаи
	if len(data) == 0 || k == 0 {
		return out
	}
	out = append(out, data[index])
	i := 1
	for i <= k {
		if index-i < 0 || index+i >= len(data) {
			break
		}
		if index == len(data)-1 && index-i >= 0 {
			out = append(out, data[index-i])
			i++
			continue
		}
		if index == 0 && index+i < len(data) {
			out = append(out, data[index+i])
			i++
			continue
		}
		if data[index-i] < data[index+1] {
			out = append(out, data[index-1])
		} else {
			out = append(out, data[index+1])
		}
		i++
		continue
	}
	return out
}
