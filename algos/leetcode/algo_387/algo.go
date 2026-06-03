package algo_387

func FirstUniqueChrAlgo387(in string) int {
	defSize := 24
	defOut := -1
	uniq := make(map[uint8]int, defSize)
	filter := make(map[uint8]int, defSize/3)

	for i := 0; i < len(in); i++ {
		ch := in[i]

		_, ok := uniq[ch]
		if !ok {
			uniq[ch] = i
		}

		_, ok = filter[ch]
		if !ok {
			filter[ch] = i
		} else {
			delete(filter, ch)
		}
	}
	if len(filter) == 0 {
		return defOut
	}
	out := len(in) - 1
	for ch, pos := range filter {
		firstPos := uniq[ch]
		if pos == firstPos && pos < out {
			out = pos
		}
	}

	return out
}
