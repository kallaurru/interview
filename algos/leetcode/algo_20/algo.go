package algo_20

func IsValidAlgo20(in string) bool {
	converter := func(val byte) int {
		switch val {
		default:
			return 1000
		case '(':
			return -1
		case '{':
			return -2
		case '[':
			return -3
		case ')':
			return 1
		case '}':
			return 2
		case ']':
			return 3
		}
	}
	state := 0
	stack := make([]int, 0, len(in)/2)
	// проверка если строка начинается с закрывающей скобки, то выходим
	if converter(in[0]) > 0 {
		return false
	}

	for i := 0; i < len(in); i++ {
		val := converter(in[i])
		if state > state+val {
			// добавили открывающую скобку это всегда хорошо
			stack = append(stack, val)
			state += val
			continue
		}
		last := stack[len(stack)-1]
		if val+last != 0 {
			return false
		}
		stack = stack[:len(stack)-1]
		state += val
	}

	return state == 0
}
