package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var s, n int
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &s)
	if err != nil {
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(2)
	}
	shards := make(map[uint32]int, s)
	for i := 0; i < n; i++ {
		var key string

		_, err = fmt.Fscan(r, &key)
		if err != nil {
			os.Exit(33)
		}
		sum := uint32(0)
		for _, b := range key {
			sum += uint32(b)
		}
		hash := sum % uint32(s)
		shards[hash] += 1
	}
	mn := shards[uint32(0)]
	mx := shards[uint32(0)]
	for i := 0; i < s; i++ {
		val := shards[uint32(i)]
		fmt.Printf("%d %d\n", i, val)
		if val < mn {
			mn = val
		}
		if val > mx {
			mx = val
		}
	}

	fmt.Printf("%s %d\n", "imbalance", mx-mn)
}
