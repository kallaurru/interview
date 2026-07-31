package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	p := []int{50, 95, 99}
	nr := func(n, p int, data []int) int {
		res := (p*n + 99) / 100
		if res < 1 {
			return 1
		}
		return res
	}
	sum := func(in []int) int {
		sum := 0
		for _, val := range in {
			sum += val
		}

		return sum
	}

	mean := func(sum, n int) int {
		res := float64(sum) / float64(n)
		res = math.Floor(res)

		return int(res)
	}
	var count int
	_, err := fmt.Fscan(r, &count)
	if err != nil {
		os.Exit(1)
	}
	if count <= 0 {
		os.Exit(0)
	}
	vals := make([]int, 0, count)
	for i := 0; i < count; i++ {
		var num int
		_, err = fmt.Fscan(r, &num)
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			return
		}
		vals = append(vals, num)
	}

	sort.Ints(vals)
	m := mean(sum(vals), count)
	w.WriteString(fmt.Sprintf("mean=%d\n", m))

	for _, item := range p {
		rr := nr(count, item, vals)
		res := vals[rr-1]
		w.WriteString(fmt.Sprintf("p%d=%d\n", item, res))
	}

	w.Flush()
}
