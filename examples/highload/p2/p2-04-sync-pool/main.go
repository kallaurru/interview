package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(1)
	}

	if n <= 0 {
		fmt.Printf("%d %d %.1f\n", 0, 0, 0.0)
		os.Exit(0)
	}
	countMiss, countHit, countGet := 0, 0, 0
	pool := 0

	for i := 0; i < n; i++ {
		var event string

		_, err = fmt.Fscan(r, &event)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		dispatch(event, &pool, &countMiss, &countHit, &countGet)
	}
	if countGet == 0 {
		countGet = 1
	}
	fmt.Printf("%d %d %.1f\n", countMiss, countHit, float64(countHit)/float64(countGet)*100)
}

func dispatch(op string, pool, countMiss, countHit, countGet *int) {
	switch op {
	case "GET":
		*countGet++
		if *pool == 0 {
			*countMiss++
		} else {
			*countHit++
			if *pool > 0 {
				*pool--
			}
		}
	case "PUT":
		*pool++
	case "GC":
		*pool = 0
	}
}
