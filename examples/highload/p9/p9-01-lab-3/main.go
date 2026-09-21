package main

import "fmt"

func main() {
	w := 8
	count := 1600000
	part := 200000

	result := make([]int, count)
	for i := 0; i < w; i++ {
		worker(i, part, result)
	}
	res := 0
	for _, val := range result {
		res += val
	}

	fmt.Println(res)
}

func worker(idx, part int, result []int) {
	start := 0
	end := 0
	if idx == 0 {
		start = 0
	} else {
		start = idx * part
	}
	end = start + part - 1

	for i := start; i <= end; i++ {
		result[i] = 1
	}
}
