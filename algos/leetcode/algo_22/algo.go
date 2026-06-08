package algo_22

func GenerateParenthesisAlgo22(n int) []string {
	itemL := []byte{'('}
	itemR := []byte{')'}
	itemFull := []byte{'(', ')'}

	// once object
	genUC1 := func(n int) string {
		useCase := make([]byte, 0, n*2)
		for i := 0; i < n; i++ {
			useCase = append(useCase, itemFull...)
		}

		return string(useCase)
	}

	// deep once object
	genUC2 := func(n int) string {
		useCase := make([]byte, 0, n*2)
		for i := 0; i < n; i++ {
			useCase = append(useCase, itemL...)
		}
		for i := 0; i < n; i++ {
			useCase = append(useCase, itemR...)
		}

		return string(useCase)
	}

	// n >= 2 internal n-1 objects
	genUC3 := func(n int) string {
		if n < 2 {
			return genUC1(n)
		}
		useCase := make([]byte, 0, n*2)
		useCase = append(useCase, itemL...)
		for i := 0; i < n-1; i++ {
			useCase = append(useCase, itemFull...)
		}
		useCase = append(useCase, itemR...)

		return string(useCase)
	}

	genUC4 := func(n, k int) []string {
		out := make([]string, 0, n*2*2)

		ext := n - k - 1
		introObjects := make([][]byte, 0, k*2)
		extraObjects := make([][]byte, 0, ext*2)

		// сначала внешние потом внутренние
		for ext > 0 {
			introObjects = append(introObjects, itemFull)
			ext--
		}
		extraObjects = append(extraObjects, itemL)
		for k > 0 {
			extraObjects = append(extraObjects, itemFull)
			k--
		}
		extraObjects = append(extraObjects, itemR)
		for i := 0; i < 2; i++ {
			tmp := make([]byte, n*2)
			for _, val := range introObjects {
				tmp = append(tmp, val...)
			}
			for _, val := range extraObjects {
				tmp = append(tmp, val...)
			}
			out = append(out, string(tmp))
		}

		return out
	}

	out := make([]string, 0, n*2*2)
	out = append(out, genUC1(n))
	if n == 1 {
		return out
	}

	out = append(out, genUC2(n))
	if n == 2 {
		return out
	}

	out = append(out, genUC3(n))
	// gen n-k-1 ext objects - k internal obj
	k := n - 2
	for k > 0 {
		out = append(out, genUC4(n, k)...)
		k--
	}

	return out
}
