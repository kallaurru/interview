package algo_5

func MaxPalindromeAlgo5(in string) string {
	var (
		start    = 0
		end      = 0
		maxL     = 1
		l, r     = 0, 0
		emptyOut = ""
	)
	source := []byte(in)
	// min len polyndrome = 2
	if len(source) == 1 {
		return emptyOut
	}
	for i := 0; i < len(source); i++ {
		if i+1 < len(source) && source[i] == source[i+1] {
			l, r = evenPalindrome(i, source)
		} else {
			l, r = oddPalindrome(i, source)
		}
		if r-l+1 > maxL {
			start = l
			end = r
			maxL = r - l + 1
		}
	}
	if maxL == 1 {
		return emptyOut
	}

	return string(source[start : end+1])
}

// запускаем если in[idx]==in[idx+1]. Проверка должна быть выше
func evenPalindrome(idx int, in []byte) (int, int) {
	// минимальный четный полиндром вначале
	if idx == 0 && len(in) > 1 && in[idx] == in[idx+1] {
		return 0, 1
	}
	if idx == len(in)-1 && in[idx-1] == in[idx] {
		return len(in) - 2, len(in) - 1
	}
	// случаи четных полиндромов с центром idx > 0
	length := 0
	for idx-length >= 0 && idx+1+length < len(in) && in[idx-length] == in[idx+1+length] {
		length++
	}

	return idx - length + 1, idx + length
}

func oddPalindrome(idx int, in []byte) (int, int) {
	if idx == 0 {
		return 0, 0
	}
	if idx == 1 && len(in) > 2 && in[0] == in[2] {
		return 0, 2
	}

	length := 0
	for idx-length >= 0 && idx+length < len(in) && in[idx-length] == in[idx+length] {
		length++
	}
	// в четном такой ситуации нет
	if length < 2 {
		return idx, idx
	}
	return idx - length + 1, idx + length - 1
}
