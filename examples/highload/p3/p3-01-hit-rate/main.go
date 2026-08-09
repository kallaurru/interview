package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var tHit, tMiss, n int

	_, err := fmt.Fscan(r, &tHit, &tMiss)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	if tHit >= tMiss {
		os.Exit(11)
	}
	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}
	h, m := 0, 0
	hr := float64(0) // .0000
	l := float64(0)  // .00
	cache := make(map[string]struct{}, n)

	//загрузили предыдущие значения
	for i := 0; i < n; i++ {
		// разбираем линию на запчасти
		var key string

		_, err = fmt.Fscan(r, &key)
		if err != nil {
			fmt.Printf("Ошибка чтения числа #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		_, ok := cache[key]
		if ok {
			h += 1
			continue
		}
		cache[key] = struct{}{}
		m += 1
	}
	hr = float64(h) / float64(h+m)
	l = float64(h*tHit+m*tMiss) / float64(h+m)

	fmt.Printf("hits=%d misses=%d\n", h, m)
	fmt.Printf("hit_rate=%.4f\n", hr)
	fmt.Printf("avg_latency=%.2f\n", l)
}
