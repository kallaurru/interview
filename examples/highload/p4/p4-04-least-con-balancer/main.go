package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var k, n int

	_, err := fmt.Fscan(r, &k, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	counter := make(map[int][]int64, k)
	for i := 0; i < n; i++ {
		var a int64
		var d int64

		_, err = fmt.Fscan(r, &a, &d)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}

		fmt.Printf("%d\n", balance(counter, k, a, d))
	}
}

func balance(counter map[int][]int64, k int, a, d int64) int {
	if k == 1 {
		return 0 // если сервер 1 его всегда отдаем
	}
	minIdx := 0
	minCon := math.MaxInt64

	for i := 0; i < k; i++ {
		conn := counter[i]
		if conn == nil {
			conn = make([]int64, 0, 8)
			counter[i] = conn
		}
		if len(conn) == 0 && len(conn) < minCon {
			// нечего чистить
			minIdx = i
			minCon = 0
			continue
		}
		// чистим
		end := 0
		for ii := 0; ii < len(conn); ii++ {
			if conn[ii] <= a {
				end++
			}
		}
		if end > 0 {
			conn = conn[end:]
			counter[i] = conn
		}
		if len(conn) < minCon {
			minCon = len(conn)
			minIdx = i
		}
	}
	// добавим соединения
	addConnection(counter, minIdx, a, d)
	return minIdx
}

func addConnection(counter map[int][]int64, idx int, a, d int64) {
	conn := counter[idx]
	conn = append(conn, a+d)
	counter[idx] = conn
}
