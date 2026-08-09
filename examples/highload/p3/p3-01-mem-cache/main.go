package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	const mib = 1048576
	var (
		n                                    int
		bKey, bVal, bOverhead, limMib, total int64
	)

	_, err := fmt.Fscan(r, &n, &bKey, &bVal, &bOverhead, &limMib)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	fits := "no"
	totalMib := float64(0)

	total = int64(n) * (bKey + bVal + bOverhead)
	totalMib = float64(total) / float64(mib)
	if total <= limMib*int64(mib) {
		fits = "yes"
	}

	fmt.Printf("total_mib=%.2f\n", totalMib)
	fmt.Printf("fits=%s\n", fits)
}
