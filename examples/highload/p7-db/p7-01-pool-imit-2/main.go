package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		n, pool, busy int
		w, maxw, sumw int64
	)

	_, err := fmt.Fscan(r, &pool, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	// todo
	/*	data := [][]int64{
			{0, 10},
			{0, 10},
			{0, 10},
			{0, 10},
		}
	*/
	timing := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		var (
			a, d int64
		)
		_, err = fmt.Fscan(r, &a, &d)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}

		// todo
		/*		item := data[i]
				a, d  = item[0], item[1]
		*/
		s := a
		busy++
		if pool-busy < 0 {
			s = timing[len(timing)-1]
			old := timing[:len(timing)-1]
			timing = old
			busy--
		}

		end := s + d
		timing = append(timing, end)
		sort.Slice(timing, func(i, j int) bool {
			return timing[i] > timing[j]
		})
		wVal := s - a
		if wVal > 0 {
			w++
			sumw += wVal
		}
		if wVal > maxw {
			maxw = wVal
		}
	}
	fmt.Printf("%d %d %d\n", w, maxw, sumw)
}
