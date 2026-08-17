package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var t, q, n int

	_, err := fmt.Fscan(r, &t, &q)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	var wait, free, start int64

	for i := 0; i < n; i++ {
		var tNow int64

		_, err = fmt.Fscan(r, &tNow)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}
		start = max(tNow, free)
		wait = int64(q) * int64(t)
		if start-tNow <= wait {
			fmt.Printf("%d\n", start)
			free = start + int64(t)
			continue
		}
		// не пропускаем запрос
		fmt.Printf("%s\n", "DROP")
	}

}
