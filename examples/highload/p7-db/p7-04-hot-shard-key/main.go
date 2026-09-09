package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var sh, n int

	_, err := fmt.Fscan(r, &sh, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	counters := make(map[uint32]int, sh)
	for i := 0; i < n; i++ {
		var key string

		_, err = fmt.Fscan(r, &key)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		shKey := fnv32a(key) % uint32(sh)
		counters[shKey] += 1
	}
	hasHotKey := false
	for i := 0; i < sh; i++ {
		c, ok := counters[uint32(i)]
		if !ok {
			continue
		}
		if c*sh > 2*n {
			hasHotKey = true
			fmt.Printf("%d %d\n", i, c)
		}
	}
	if !hasHotKey {
		fmt.Printf("%s\n", "OK")
	}
}

func fnv32a(s string) uint32 {
	h := uint32(2166136261)       // offset basis
	for i := 0; i < len(s); i++ { // перебор БАЙТОВ строки
		h ^= uint32(s[i]) // сначала XOR…
		h *= 16777619     // …потом умножение (uint32 сам даёт mod 2^32)
	}
	return h
}
