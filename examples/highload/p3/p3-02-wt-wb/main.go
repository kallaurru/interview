package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var wT, wB, s, n int

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	dirty := make(map[string]struct{}, n)

	for i := 0; i <= n; i++ {
		var op, key string

		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(11)
		}
		parts := strings.Fields(line)
		if len(parts) == 1 {
			key = ""
			op = parts[0]
		} else if len(parts) > 1 {
			op = parts[0]
			key = parts[1]
		}

		switch op {
		default:
			continue
		case "SET":
			wT++
			_, ok := dirty[key]
			if !ok {
				dirty[key] = struct{}{}
			}
		case "FLUSH":
			wB += len(dirty)
			dirty = make(map[string]struct{}, n)
		}
	}
	// финальный сброс для WB
	wB += len(dirty)
	dirty = make(map[string]struct{}, n)
	s = wT - wB
	fmt.Printf("write_through=%d write_behind=%d saved=%d\n", wT, wB, s)
}
