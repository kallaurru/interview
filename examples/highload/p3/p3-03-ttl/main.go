package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stat struct {
	Hits int
	Miss int
	Cold int
}

func main() {
	r := bufio.NewReader(os.Stdin)

	var n int

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	if n < 1 {
		os.Exit(0) // нет работы у кэша
	}

	cache := make(map[string]int, n)
	stat := Stat{}
	for i := 0; i <= n; i++ {
		// разбираем линию на запчасти
		var t, ttl int
		var op, key string

		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(11)
		}

		parts := strings.Fields(line)
		if len(parts) > 2 {
			op = parts[1]
			key = parts[2]
			t, err = strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("Ошибка конвертации в строке #%d: %v\n", i+1, err)
				os.Exit(22)
			}
		}
		if len(parts) > 3 {
			ttl, err = strconv.Atoi(parts[3])
			if err != nil {
				fmt.Printf("Ошибка конвертации в строке #%d: %v\n", i+1, err)
				os.Exit(33)
			}
		}

		// todo del after test
		fmt.Printf("T- %d, OP- %s KEY- %s TTL - %d\n", t, op, key, ttl)

		switch op {
		default:
			continue
		case "GET":
			t1, ok := cache[key]
			if !ok {
				stat.Cold++ // холодный промах ключа не было
				continue
			}
			if t < t1 {
				stat.Hits++
				continue
			}
			// ленивое удаление
			stat.Miss++ // промах по истечении ttl
			delete(cache, key)
		case "PUT":
			cache[key] = t + ttl
		}
	}
	fmt.Printf("hits=%d misses_cold=%d misses_expired=%d\n", stat.Hits, stat.Cold, stat.Miss)
}
