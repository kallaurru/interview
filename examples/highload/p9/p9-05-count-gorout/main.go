package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Info struct {
	Name  string
	Delta int
}

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(1)
	}
	diffs := make([]Info, 0, n)
	for i := 0; i < n; i++ {
		var (
			c1, c2 int
			name   string
		)

		_, err = fmt.Fscan(r, &name, &c1, &c2)
		if err != nil {
			os.Exit(2)
		}

		if c2-c1 > 0 {
			diffs = append(diffs, Info{
				Name:  name,
				Delta: c2 - c1,
			})
		}
	}
	if len(diffs) == 0 {
		fmt.Printf("%s\n", "no leak")
		os.Exit(0)
	}
	sort.Slice(diffs, func(i, j int) bool {
		if diffs[i].Delta == diffs[j].Delta {
			return diffs[i].Name < diffs[j].Name
		}
		return diffs[i].Delta > diffs[j].Delta
	})
	sum := 0
	for _, info := range diffs {
		fmt.Printf("%s +%d\n", info.Name, info.Delta)
		sum += info.Delta
	}
	fmt.Printf("TOTAL +%d\n", sum)
}
