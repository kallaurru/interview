package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fnv1a32(s string) uint32 {
	h := uint32(2166136261)
	for i := 0; i < len(s); i++ {
		h ^= uint32(s[i])
		h *= 16777619
	}
	return h
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	if !sc.Scan() {
		return
	}

	parts := strings.Fields(sc.Text())
	if len(parts) < 2 {
		return
	}

	N, _ := strconv.Atoi(parts[0])
	M, _ := strconv.Atoi(parts[1])

	n32 := uint32(N)
	m32 := uint32(M)

	total := 0
	moved := 0

	for sc.Scan() {
		key := sc.Text()
		total++

		h := fnv1a32(key)
		if h%n32 != h%m32 {
			moved++
		}
	}

	percent := 0
	if total > 0 {
		percent = moved * 100 / total
	}

	fmt.Printf("total=%d\n", total)
	fmt.Printf("moved=%d\n", moved)
	fmt.Printf("percent=%d\n", percent)
}
