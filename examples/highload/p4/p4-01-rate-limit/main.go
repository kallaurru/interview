package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var k, n int

	_, err := fmt.Fscan(r, &k)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	clients := make([]string, 0, k)
	quotas := make(map[string]int, k)
	// загружаем квоты
	for i := 0; i < k; i++ {
		var name string
		var quo int

		_, err = fmt.Fscan(r, &name, &quo)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(2)
		}
		_, ok := quotas[name]
		if !ok {
			clients = append(clients, name)
		}
		quotas[name] = quo

	}
	counterAllowed := make(map[string]int, k)
	counterDenied := make(map[string]int, k)

	// грузим количество событий
	_, err = fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	// обрабатываем поток событий

	for i := 0; i < n; i++ {
		// разбираем линию на запчасти
		var name = ""
		_, err = fmt.Fscan(r, &name)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		quota := quotas[name]
		if quota > 0 {
			quotas[name] -= 1
			counterAllowed[name] += 1
			continue
		}
		counterDenied[name] += 1
	}
	for _, client := range clients {
		allowed := counterAllowed[client]
		denied := counterDenied[client]
		fmt.Printf("%s %d %d\n", client, allowed, denied)
	}
}
