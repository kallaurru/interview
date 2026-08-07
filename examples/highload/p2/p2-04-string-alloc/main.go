package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		n          int
		dummyLen   int64
		dummyCap   int64
		countBCap  int64
		builderCap int64
		builderLen int64
	)

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(1)
	}

	if n <= 0 {
		os.Exit(11)
	}

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(r, &numbers[i])
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		dummyLen += int64(numbers[i])
		dummyCap += dummyLen
		builderLen += int64(numbers[i])
		if builderLen > builderCap {
			for builderLen > builderCap {
				if builderCap == 0 {
					builderCap++
					continue
				}
				builderCap *= 2
			}
			countBCap += builderCap
		}
	}

	fmt.Printf("%d %d %.1f\n", dummyCap, countBCap, float64(dummyCap)/float64(countBCap))
}
