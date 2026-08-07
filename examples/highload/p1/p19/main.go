package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	const basis uint32 = 2166136261
	const constH1 uint32 = 16777619
	var s, n int

	_, err := fmt.Fscan(r, &s, &n)
	if err != nil {
		os.Exit(1)
	}

	keys := make([]int, s, s)

	for i := 0; i < n; i++ {
		var val []byte
		_, err = fmt.Fscan(r, &val)

		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		hash := fnv(val, basis, constH1)
		sID := int(hash % uint32(s))
		keys[sID] += 1
		fmt.Printf("%s %d\n", string(val), sID)
	}
	buf := new(bytes.Buffer)
	for _, val := range keys {
		buf.WriteString(fmt.Sprintf(" %d", val))
	}
	str := buf.String()
	str = strings.Trim(str, " ")

	fmt.Printf("%s\n", str)
}

func fnv(key []byte, basis, h1 uint32) uint32 {
	var out uint32
	out = basis
	for _, b := range key {
		out ^= uint32(b)
		out = (out * h1) % (1<<32 - 1)
	}

	return out
}
