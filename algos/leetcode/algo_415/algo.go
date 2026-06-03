package algo_415

func AddStringsAlgo415(num1, num2 string) string {
	sync := func(i, diff int) int {
		return i - diff
	}
	forward := map[byte]int{
		'0': 0,
		'1': 1,
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
	}
	reverse := map[int]byte{
		0: '0',
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
	}
	if len(num1) > len(num2) {
		// если разные длины num1 всегда делаем коротким
		num1, num2 = num2, num1
	}
	out := make([]byte, len(num2)+1)
	rem := 0
	diff := len(num2) - len(num1)
	// num2 всегда длиннее
	for i := len(num2) - 1; i >= 0; i-- {
		short := 0
		if sync(i, diff) >= 0 {
			short = forward[num1[sync(i, diff)]]
		}
		val := short + forward[num2[i]] + rem
		if val > 10 {
			rem = 1
			out[i+1] = reverse[val-10]
			continue
		}
		rem = 0
		out[i+1] = reverse[val]
	}

	if rem > 0 {
		out[0] = reverse[rem]
		return string(out)
	}

	return string(out[1:])
}
