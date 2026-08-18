package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var k, w, n, m int

	_, err := fmt.Fscan(r, &k, &n, &w)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &m)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	var local, central int
	gCalc := globCalc(n, w)
	lCalc := localCalc(n, w, k)
	for i := 0; i < m; i++ {
		var node int
		var t int

		_, err = fmt.Fscan(r, &t, &node)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}

		central += gCalc(t)
		local += lCalc(t, node)
	}

	fmt.Printf("local %d\n", local)
	fmt.Printf("central %d\n", central)

}
func localCalc(n, w, k int) func(t, node int) int {
	nodes := make(map[int]map[int]int, k)

	return func(t, node int) int {
		if n < k {
			return 0 // нет квот
		}
		limPerNode := n / k
		winN := t / w
		counters, ok := nodes[node]
		if !ok {
			counters = make(map[int]int, n)
			counters[winN] = 1
			nodes[node] = counters
			return 1
		}
		count, ok := counters[winN]
		if !ok {
			counters[winN] += 1
			nodes[node] = counters
			return 1
		}
		if count < limPerNode {
			counters[winN] += 1
			nodes[node] = counters
			return 1
		}
		return 0
	}
}

func globCalc(n, w int) func(t int) int {
	counter := make(map[int]int, 8)

	return func(t int) int {
		winN := t / w
		count, ok := counter[winN]
		if !ok {
			if n > 0 {
				counter[winN] = 1
				return 1
			}
			// если количество запросов 0 на окно W
			return 0
		}
		if count+1 <= n {
			counter[winN] = count + 1
			return 1
		}

		return 0
	}
}
