package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var l, w, n int

	_, err := fmt.Fscan(r, &l, &w)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	clients := make(map[string][]int64, (n+1)/2)
	for i := 0; i < n; i++ {
		var (
			tNow   int64
			client string
		)

		_, err = fmt.Fscan(r, &client, &tNow)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}
		q := clients[client]
		// если массива еще нет
		if q == nil || len(q) == 0 {
			q = make([]int64, 0, 8)
			q = append(q, tNow)
			clients[client] = q
			fmt.Printf("%s\n", "ALLOW")
			continue
		}
		ii := 0
		for ii < len(q) && q[ii] <= tNow-int64(w) {
			ii++
		}
		q = q[ii:]
		if len(q) < l {
			q = append(q, tNow)
			fmt.Printf("%s\n", "ALLOW")
		} else {
			// не пропускаем запрос
			fmt.Printf("%s\n", "DENY")
		}
		clients[client] = q
	}

}
