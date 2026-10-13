package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	// Read first line: quantiles
	if !sc.Scan() {
		return
	}
	qStrs := strings.Fields(sc.Text())
	quantiles := make([]float64, len(qStrs))
	for i, s := range qStrs {
		q, _ := strconv.ParseFloat(s, 64)
		quantiles[i] = q
	}

	type bucket struct {
		le    float64
		cum   float64
		isInf bool
	}

	var buckets []bucket
	var total float64

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}
		leStr := parts[0]
		cumStr := parts[1]

		cum, _ := strconv.ParseFloat(cumStr, 64)

		var le float64
		var isInf bool
		if leStr == "+Inf" {
			le = math.Inf(1)
			isInf = true
			total = cum
		} else {
			le, _ = strconv.ParseFloat(leStr, 64)
		}
		buckets = append(buckets, bucket{le: le, cum: cum, isInf: isInf})
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for i, q := range quantiles {
		rank := q * total

		// Find first bucket with cum >= rank
		found := -1
		for j, b := range buckets {
			if b.cum >= rank {
				found = j
				break
			}
		}

		var answer float64

		if found == -1 {
			// Fallback (should not happen since last bucket cum = total >= rank)
			answer = 0
		} else {
			b := buckets[found]
			var lower, prev float64
			if found == 0 {
				lower = 0
				prev = 0
			} else {
				lower = buckets[found-1].le
				prev = buckets[found-1].cum
			}

			if b.isInf {
				answer = lower
			} else {
				denom := b.cum - prev
				if denom == 0 {
					answer = lower
				} else {
					answer = lower + (rank-prev)/denom*(b.le-lower)
				}
			}
		}

		fmt.Fprintf(out, "%s %.3f\n", qStrs[i], answer)
	}
}
