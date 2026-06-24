package algo_71

func SimplifyPathAlgo71(path string) string {
	stack := make([]string, 0, 8)
	add := func(pie string, stack []string) []string {
		switch pie {
		default:
			stack = append(stack, pie)
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
		}

		return stack
	}

	l := 0
	for i := 0; i < len(path); i++ {
		if path[i] != '/' {
			continue
		}
		if i-l == 0 {
			l++
			continue
		}
		pie := path[l:i]
		stack = add(pie, stack)
		l = i + 1
	}
	if path[len(path)-1] != '/' {
		stack = add(path[l:], stack)
	}
	out := make([]byte, 0, 8)
	out = append(out, '/')
	for _, item := range stack {
		out = append(out, []byte(item)...)
		out = append(out, '/')
	}
	out = out[:len(out)-1]

	return string(out)
}
