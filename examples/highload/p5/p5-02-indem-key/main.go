package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		n int
	)

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	uniq, duplicate := 0, 0
	total := int64(0)
	keys := make(map[string]struct{}, n)

	for i := 0; i < n; i++ {
		var (
			key    string
			amount int64
		)

		_, err = fmt.Fscan(r, &key, &amount)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		_, ok := keys[key]

		keyword := "DUPLICATE"
		if !ok {
			total += amount
			keys[key] = struct{}{}
			uniq++
			keyword = "APPLIED"
		} else {
			duplicate++
		}

		fmt.Printf("%s %s %d\n", keyword, key, total)
	}

	fmt.Printf("TOTAL %d UNIQUE %d DUPLICATES %d\n", total, uniq, duplicate)
}
