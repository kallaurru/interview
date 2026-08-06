package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	const prefix = 8

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(1)
	}
	if n == 0 {
		os.Exit(0) // не работы
	}
	oneAlloc := n * prefix
	count, multiAlloc := calcMultiAlloc(n, prefix)
	diff := float64(multiAlloc) / float64(oneAlloc)

	fmt.Printf("%d %d %d %.2f\n", count, multiAlloc, oneAlloc, diff)
}

func calcMultiAlloc(n, prefix int) (int, uint64) {

	length, capa, alloc, countAlloc := 1, 1, 1, 1
	for i := 1; i < n; i++ {
		if length == capa {
			capa *= 2
			alloc += capa
			countAlloc++
		}
		length += 1
	}

	alloc *= prefix

	return countAlloc, uint64(alloc)
}
