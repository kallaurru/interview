package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Bucket struct {
	leStr string
	leVal float64
	isInf bool
	sum   int64
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	buckets := make(map[string]*Bucket)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 3 {
			continue
		}
		// pod := parts[0] // not needed
		leStr := parts[1]
		cumStr := parts[2]

		cum, err := strconv.ParseInt(cumStr, 10, 64)
		if err != nil {
			continue
		}

		var leVal float64
		var isInf bool
		if leStr == "+Inf" {
			leVal = math.Inf(1)
			isInf = true
		} else {
			leVal, err = strconv.ParseFloat(leStr, 64)
			if err != nil {
				continue
			}
		}

		if b, ok := buckets[leStr]; ok {
			b.sum += cum
		} else {
			buckets[leStr] = &Bucket{
				leStr: leStr,
				leVal: leVal,
				isInf: isInf,
				sum:   cum,
			}
		}
	}

	// Collect and sort
	var list []*Bucket
	for _, b := range buckets {
		list = append(list, b)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].leVal < list[j].leVal
	})

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var total int64
	for _, b := range list {
		fmt.Fprintf(out, "%s %d\n", b.leStr, b.sum)
		if b.isInf {
			total = b.sum
		}
	}

	// Compute p95
	rank := 0.95 * float64(total)

	found := -1
	for i, b := range list {
		if float64(b.sum) >= rank {
			found = i
			break
		}
	}

	var answer float64
	if found == -1 {
		answer = 0
	} else {
		b := list[found]
		var lower, prev float64
		if found == 0 {
			lower = 0
			prev = 0
		} else {
			lower = list[found-1].leVal
			prev = float64(list[found-1].sum)
		}

		if b.isInf {
			answer = lower
		} else {
			denom := float64(b.sum) - prev
			if denom == 0 {
				answer = lower
			} else {
				answer = lower + (rank-prev)/denom*(b.leVal-lower)
			}
		}
	}

	fmt.Fprintf(out, "p95 %.3f\n", answer)
}
