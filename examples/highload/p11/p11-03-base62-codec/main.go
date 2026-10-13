package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// encode converts a non-negative integer to its base62 representation.
func encode(n uint64) string {
	if n == 0 {
		return "0"
	}
	var sb strings.Builder
	for n > 0 {
		rem := n % 62
		sb.WriteByte(alphabet[rem])
		n /= 62
	}
	// Reverse the string
	b := []byte(sb.String())
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

// decode converts a base62 string to its integer value.
func decode(code string) uint64 {
	var n uint64
	for i := 0; i < len(code); i++ {
		c := code[i]
		var idx uint64
		switch {
		case c >= '0' && c <= '9':
			idx = uint64(c - '0')
		case c >= 'a' && c <= 'z':
			idx = uint64(c-'a') + 10
		case c >= 'A' && c <= 'Z':
			idx = uint64(c-'A') + 36
		}
		n = n*62 + idx
	}
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		if len(parts) < 2 {
			continue
		}
		cmd := parts[0]
		arg := parts[1]

		switch cmd {
		case "enc":
			n, err := strconv.ParseUint(arg, 10, 64)
			if err != nil {
				continue // invalid input, skip
			}
			fmt.Println(encode(n))
		case "dec":
			fmt.Println(decode(arg))
		}
	}
}
