package algo_392

func IsSubsequenceAlgo392(s, t string) bool {
	sPtr := 0
	count := len(s)

	if len(s) == 0 && len(t) == 0 {
		return true
	}

	if len(s) == 0 || len(t) == 0 {
		return false
	}

	for i := 0; i < len(t); i++ {
		if s[sPtr] == t[i] {
			count--
			sPtr++
		}
		if count == 0 {
			break
		}
		if sPtr >= len(s) {
			break
		}
	}

	return count == 0
}

/*
	для альтернативного решения когда проверяется больше 100 s подстрок.
	Нужно загнать t в map[byte]int. И затем отслеживать что бы следующий
	символ s занимал большую позицию в t.
*/
