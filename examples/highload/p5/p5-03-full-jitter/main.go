package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		base, capacity, attempts int
		seed, totalSleep         int64
	)

	_, err := fmt.Fscan(r, &base, &capacity, &attempts, &seed)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	var delayVal, raw int64
	for i := 1; i <= attempts; i++ {
		delayVal, raw, seed = delay(capacity, base, i, seed)
		totalSleep += delayVal
		fmt.Printf("%d %d %d\n", i, raw, delayVal)
	}

	fmt.Printf("TOTAL_SLEEP %d\n", totalSleep)
}

func delay(cap, base, i int, seed int64) (int64, int64, int64) {
	const (
		seed1  = int64(1103515245)
		seed2  = int64(12345)
		modVal = int64(2147483648)
	)

	var (
		raw, delayVal int64
	)
	powVal := int64(1 << (i - 1))
	raw = min(int64(cap), int64(base)*powVal)
	seed = (seed*seed1 + seed2) % modVal
	delayVal = seed % (raw + 1)

	return delayVal, raw, seed
}
