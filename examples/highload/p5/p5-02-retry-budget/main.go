package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	canRetry := func(method, code string) bool {
		return method != "POST" && (code == "429" || code == "503")
	}
	r := bufio.NewReader(os.Stdin)
	var (
		n int
	)

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	// загружаем прыжки
	allAttempts := 0

	for i := 0; i < n; i++ {
		var reqId, method, r1, r2, r3 string

		_, err = fmt.Fscan(r, &reqId, &method, &r1, &r2, &r3)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}
		results := [3]string{r1, r2, r3}

		for idx, val := range results {
			allAttempts++
			if val == "OK" {
				fmt.Printf("%s %s %d\n", reqId, val, idx+1)
				goto ext
			}
			if canRetry(method, val) {
				continue
			}

			fmt.Printf("%s %s %d\n", reqId, val, idx+1)
			goto ext
		}
		// вариант когда все повторы закончились, а результата нет
		fmt.Printf("%s %s %d\n", reqId, results[len(results)-1], 3)
	ext:
	}
	fmt.Printf("AMPLIFICATION %.2f\n", float64(allAttempts)/float64(n))
}
