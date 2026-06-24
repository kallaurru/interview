package algo_56

import (
	"sort"
)

func MergeIntervalsAlgo56(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	out := make([][]int, 0, len(intervals)/2)
	out = append(out, intervals[0])

	for i := 1; i < len(intervals); i++ {
		last := out[len(out)-1]
		curr := intervals[i]
		if curr[0] <= last[1] {
			if curr[1] > last[1] {
				last[1] = curr[1]
			}
			continue
		}

		out = append(out, curr)
	}
	return out
}
