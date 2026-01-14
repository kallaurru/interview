package tnearkfori

func NearKElements(data []int, index, k int) []int {
	out := make([]int, 0, k)
	// крайние случаи
	if len(data) == 0 || k == 0 {
		return out
	}
	// если число в адресе index тоже считаем ближайшим
	out = append(out, data[index])
	k--
	l, r := index-1, index+1

	for k > 0 {
		k--
		if r > len(data)-1 {
			out = append(out, data[l])
			l--
			continue
		}
		if l < 0 {
			out = append(out, data[r])
			r++
			continue
		}
		if data[index]-data[l] < data[r]-data[index] {
			out = append(out, data[l])
			l--
			continue
		}
		out = append(out, data[r])
		r++
	}
	return out
}
