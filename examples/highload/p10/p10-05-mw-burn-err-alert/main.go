package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	if !sc.Scan() {
		return
	}
	parts := strings.Fields(sc.Text())
	if len(parts) < 4 {
		return
	}
	SLO, _ := strconv.ParseFloat(parts[0], 64)
	Ws, _ := strconv.Atoi(parts[1])
	Wl, _ := strconv.Atoi(parts[2])
	T, _ := strconv.ParseFloat(parts[3], 64)

	denom := 1.0 - SLO/100.0

	var bads []int64
	var totals []int64

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		p := strings.Fields(line)
		if len(p) < 2 {
			continue
		}
		g, _ := strconv.ParseInt(p[0], 10, 64)
		t, _ := strconv.ParseInt(p[1], 10, 64)
		bads = append(bads, t-g)
		totals = append(totals, t)
	}

	n := len(bads)

	// Prefix sums (1-indexed)
	prefBad := make([]int64, n+1)
	prefTotal := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		prefBad[i] = prefBad[i-1] + bads[i-1]
		prefTotal[i] = prefTotal[i-1] + totals[i-1]
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	anyAlert := false

	for t := 1; t <= n; t++ {
		// Short window
		sStart := t - Ws + 1
		if sStart < 1 {
			sStart = 1
		}
		sBad := prefBad[t] - prefBad[sStart-1]
		sTotal := prefTotal[t] - prefTotal[sStart-1]
		var sBurn float64
		if sTotal > 0 {
			sBurn = (float64(sBad) / float64(sTotal)) / denom
		}

		// Long window
		lStart := t - Wl + 1
		if lStart < 1 {
			lStart = 1
		}
		lBad := prefBad[t] - prefBad[lStart-1]
		lTotal := prefTotal[t] - prefTotal[lStart-1]
		var lBurn float64
		if lTotal > 0 {
			lBurn = (float64(lBad) / float64(lTotal)) / denom
		}

		if sBurn >= T && lBurn >= T {
			fmt.Fprintln(out, t)
			anyAlert = true
		}
	}

	if !anyAlert {
		fmt.Fprintln(out, "no alerts")
	}
}
