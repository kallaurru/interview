package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	var (
		dau, rr, p, u, b int
	)

	_, err := fmt.Fscan(r, &dau, &rr, &p, &u, &b)
	if err != nil {
		os.Exit(1)
	}
	avg, peak, sizemb := calculator(dau, rr, p, u, b)

	fmt.Printf("avg_qps=%d\npeak_qps=%d\nstate_mb=%d\n", avg, peak, sizemb)
}

func calculator(dau, r, p, u, b int) (int64, int64, int64) {
	const secPerDay int64 = 86400
	const memSize int64 = 1048576

	var avgQ, peakQ, stateMb int64

	avgQ = (int64(dau)*int64(r) + (secPerDay - 1)) / secPerDay
	peakQ = avgQ * int64(p)
	stateMb = (int64(u)*int64(b) + (memSize - 1)) / memSize

	return avgQ, peakQ, stateMb
}
