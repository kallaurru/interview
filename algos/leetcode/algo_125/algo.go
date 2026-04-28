package algo_125

import "strings"

func ValidPalindrome(in string) bool {
	alphaNumF := func(in string) string {
		out := make([]byte, 0, len(in))
		for i := 0; i < len(in); i++ {
			if ('a' <= in[i] && in[i] <= 'z') || ('0' <= in[i] && in[i] <= '9') {
				out = append(out, in[i])
			}
		}

		return string(out)
	}
	ff := func(in string) string {
		in = strings.ToLower(in)
		in = alphaNumF(in)
		return in
	}
	// четный abba abddba здесь центер правый символ
	evenP := func(in string, center int) bool {
		r := 0
		if in[center-1] != in[center] {
			return false
		}
		for center-r-1 >= 0 && center+r < len(in) && in[center-1-r] == in[center+r] {
			r++
		}
		return r*2 == len(in)
	}
	// нечетный bacab
	oddP := func(in string, center int) bool {
		r := 0
		for center-r >= 0 && center+r < len(in) && in[center-r] == in[center+r] {
			r++
		}
		return (r-1)*2+1 == len(in)
	}

	clearIn := ff(in)
	if clearIn == "" {
		return true // пустая строка
	}

	rem := len(clearIn) % 2
	center := len(clearIn) / 2

	if rem == 0 {
		return evenP(clearIn, center)
	}

	return oddP(clearIn, center)
}
