package main

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var c, tps, n int

	_, err := fmt.Fscan(r, &c, &tps)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	var tokens atomic.Int64
	var prevT atomic.Int64
	const beginEpoch = int64(-1)

	tokens.Store(int64(c))
	prevT.Store(beginEpoch)

	for i := 0; i < n; i++ {
		var cost int64
		var t int64

		_, err = fmt.Fscan(r, &t, &cost)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}

		if prevT.Load() == beginEpoch {
			prevT.Store(t)
			if tokens.Load() >= cost {
				printRes(true)
				tokens.Store(maxCustom(0, tokens.Load()-cost))
				continue
			}
			printRes(false)
			continue
		}
		diff := (t - prevT.Load()) * int64(tps)

		tok := minCustom(int64(c), tokens.Load()+diff)
		prevT.Store(t)
		if tok >= cost {
			tokens.Store(maxCustom(0, tok-cost))
			printRes(true)
			continue
		}
		printRes(false)
	}

}

func printRes(res bool) {
	if res {
		fmt.Printf("%s\n", "ALLOW")
	} else {
		fmt.Printf("%s\n", "DENY")
	}
}

func minCustom(one, two int64) int64 {
	if one < two {
		return one
	}

	return two
}

func maxCustom(one, two int64) int64 {
	if one > two {
		return one
	}

	return two
}
