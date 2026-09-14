package main

import (
	"bufio"
	"fmt"
	"os"
)

type Item struct {
	Body   string
	Moment int
}

func main() {
	var n, ttl int
	r := bufio.NewReader(os.Stdin)

	_, err := fmt.Fscan(r, &ttl)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}
	cache := make(map[string]Item, (n+1)/2)
	for i := 0; i < n; i++ {
		var (
			ts        int
			key, body string
		)
		_, err = fmt.Fscan(r, &ts, &key, &body)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(3)
		}
		item, ok := cache[key]
		if !ok {
			item = Item{
				Body:   body,
				Moment: ts,
			}
			cache[key] = item
			fmt.Printf("%s\n", "PROCESS")
			continue
		}
		// запись истекла нужно обновить
		if ts-item.Moment >= ttl {
			item.Moment = ts
			item.Body = body
			cache[key] = item
			fmt.Printf("%s\n", "PROCESS")
			continue
		}

		if ts-item.Moment < ttl && body == item.Body {
			fmt.Printf("%s\n", "DUPLICATE")
			continue
		}

		if ts-item.Moment < ttl && body != item.Body {
			fmt.Printf("%s\n", "CONFLICT")
			continue
		}
	}

}
