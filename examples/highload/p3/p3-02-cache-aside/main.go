package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var hits, miss, dbR, dbW, n int

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	cache := make(map[string]struct{}, n)

	//загрузили предыдущие значения
	for i := 0; i < n; i++ {
		// разбираем линию на запчасти
		var op, key string

		_, err = fmt.Fscan(r, &op, &key)
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		switch op {
		default:
			continue
		case "GET":
			_, ok := cache[key]
			if ok {
				hits++
				continue
			}
			miss++
			dbR++
			cache[key] = struct{}{}
		case "SET":
			dbW++
			delete(cache, key)
		}
	}

	fmt.Printf("hits=%d misses=%d db_reads=%d db_writes=%d\n", hits, miss, dbR, dbW)
}
