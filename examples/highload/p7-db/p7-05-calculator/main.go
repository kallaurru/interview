package main

import (
	"bufio"
	"fmt"
	"os"
)

type CalcResult struct {
	AvgRps    int64
	PeakRps   int64
	WritesDay int64
	StorGb    int64
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var dau, rpd, rr, w, rOw, repl, ret int64

	_, err := fmt.Fscan(r, &dau, &rpd, &rr, &w, &rOw, &repl, &ret)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	res := calculator(dau, rpd, rr, w, rOw, repl, ret)

	fmt.Printf("%d %d %d %d\n", res.AvgRps, res.PeakRps, res.WritesDay, res.StorGb)
}

func calculator(dau, rpd, rr, w, rOw, repl, ret int64) CalcResult {
	const shift = 1073741824

	res := CalcResult{}
	total := dau * rpd
	res.AvgRps = (total + 86400 - 1) / 86400
	res.PeakRps = res.AvgRps * 3
	res.WritesDay = total * w / (rr + w)
	if res.WritesDay == 0 {
		res.StorGb = 0
	} else {
		res.StorGb = ((res.WritesDay * rOw * repl * ret) + shift - 1) / shift
	}

	return res
}
