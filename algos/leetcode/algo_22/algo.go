package algo_22

func GenerateParenthesisAlgo22(n int) []string {
	itemL := []byte{'('}
	itemR := []byte{')'}
	itemFull := []byte{'(', ')'}

	out := make([]string, 0, n*2-1)

	extBuilder := func(n int) []byte {
		items := make([]byte, 0, n*2)
		for i := 0; i < n; i++ {
			items = append(items, itemFull...)
		}
		return items
	}

	deepInternalBuilder := func(n int) []byte {
		if n < 2 {
			return []byte{}
		}
		items := make([]byte, 0, n*2)
		for i := 0; i < n-1; i++ {
			items = append(items, itemL...)
		}
		items = append(items, itemFull...)
		for i := 0; i < n-1; i++ {
			items = append(items, itemR...)
		}

		return items
	}

	internalBuilder := func(n int) []byte {
		items := make([]byte, 0, n*2)
		items = append(items, itemL...)
		for i := 0; i < n; i++ {
			items = append(items, itemFull...)
		}
		items = append(items, itemR...)

		return items
	}

	item := extBuilder(n)
	out = append(out, string(item))

	if n == 1 {
		return out
	}

	item = deepInternalBuilder(n)
	out = append(out, string(item))

	if n == 2 {
		return out
	}

	item = internalBuilder(n - 1)
	out = append(out, string(item))
	maxInternal := n - 1 - 1
	for maxInternal > 0 {
		paar := make([]byte, 0, n*2*2-1)
		reversePaar := make([]byte, 0, n*2*2-1)
		paar = append(paar, internalBuilder(maxInternal)...)
		paar = append(paar, extBuilder(n-1-maxInternal)...)
		out = append(out, string(paar))
		reversePaar = append(reversePaar, extBuilder(n-1-maxInternal)...)
		reversePaar = append(reversePaar, internalBuilder(maxInternal)...)
		out = append(out, string(reversePaar))
		maxInternal--
	}

	return out
}
