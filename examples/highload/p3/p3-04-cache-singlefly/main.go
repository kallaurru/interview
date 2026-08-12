package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stat struct {
	DBCalls int
	CoalEsc int
	Hits    int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var d, n int

	_, err := fmt.Fscan(r, &d)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	cache := make(map[string]struct{}, n)
	sf := make(map[string]int, (n+1)/2)
	stat := Stat{}
	for i := 0; i < n; i++ {
		// разбираем линию на запчасти
		var (
			t   = 0
			key = ""
		)

		_, err = fmt.Fscan(r, &t, &key)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		_, ok := cache[key]
		if ok {
			stat.Hits++
			continue
		}
		tx, ok := sf[key]
		if !ok {
			// запускаем полет
			sf[key] = t + d
			stat.DBCalls++
			continue
		}
		if t >= tx {
			delete(sf, key) // убрали из полета
			cache[key] = struct{}{}
			// ключ уже в кэше
			stat.Hits++
			continue
		}
		// присоединяемся к полету
		stat.CoalEsc++
	}

	fmt.Printf("db_calls=%d coalesced=%d hits=%d naive_db_calls=%d\n",
		stat.DBCalls, stat.CoalEsc, stat.Hits, stat.DBCalls+stat.CoalEsc)
}
