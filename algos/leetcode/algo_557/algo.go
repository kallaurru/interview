package algo_557

func ReverseWordAlgo557(in string) string {
	out := make([]byte, len(in))
	space := uint8(32)
	l, r := len(in)-1, len(in)-1

	for l = len(in) - 1; l >= 0; l-- {
		out[l] = in[l]
		if in[l] != space && in[r] == space {
			// начало нового слова
			r = l
			continue
		}
		if in[l] != space {
			continue
		}
		// такой ситуации по условиям нет, но мы ее здесь опишем
		if in[l] == space && in[r] == space {
			continue
		}
		countSwap := (r - l) / 2
		for k := 0; k < countSwap; k++ {
			out[l+1+k], out[r-k] = out[r-k], out[l+1+k]
		}
		r = l
		continue
	}
	if r-0 != 0 {
		// перестановка последнего слова
		countSwap := (r - l) / 2
		for k := 0; k < countSwap; k++ {
			out[l+1+k], out[r-k] = out[r-k], out[l+1+k]
		}
	}

	return string(out)
}
