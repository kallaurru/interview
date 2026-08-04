package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var k, m int

	_, err := fmt.Fscan(r, &k, &m)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}

	names := make(map[int]string)
	oldResults := make([]int, 0, k)
	printCache := make([]string, 0, k)
	flgLim := true

	//загрузили предыдущие значения
	for i := 0; i < k; i++ {
		// разбираем линию на запчасти
		var item string
		values := make([]int, 0, m)
		for ll := 0; ll <= m; ll++ {
			_, err = fmt.Fscan(r, &item)
			if err != nil {
				fmt.Printf("Ошибка чтения числа #%d: %v\n", ll+1, err)
				os.Exit(2)
			}
			if ll == 0 {
				names[i] = item
				continue
			}
			val, err := strconv.Atoi(item)
			if err != nil {
				fmt.Printf("Ошибка чтения числа #%s idx - %d: %v\n", item, ll, err)
				os.Exit(3)
			}
			values = append(values, val)
		}
		sort.Ints(values)
		idxMed := m / 2
		oldResults = append(oldResults, values[idxMed])
	}

	// загружаем новые
	for i := k; i < k*2; i++ {
		// разбираем линию на запчасти
		var item string
		values := make([]int, 0, m)
		for ll := 0; ll <= m; ll++ {
			_, err = fmt.Fscan(r, &item)
			if err != nil {
				fmt.Printf("Ошибка чтения числа #%d: %v\n", ll+1, err)
				os.Exit(2)
			}
			if ll == 0 {
				continue
			}
			val, err := strconv.Atoi(item)
			if err != nil {
				fmt.Printf("Ошибка чтения числа #%s idx - %d: %v\n", item, ll, err)
				os.Exit(3)
			}
			values = append(values, val)
		}
		sort.Ints(values)
		idxMed := m / 2
		thIdx := i % k
		newMed := values[idxMed]
		oldMed := oldResults[thIdx]
		delta := float64(newMed-oldMed) / float64(oldMed) * 100
		if delta > 5 {
			flgLim = false
		}
		name, ok := names[thIdx]
		if ok {
			printCache = append(printCache, fmt.Sprintf("%s %+.1f%%", name, delta))
		} else {
			fmt.Printf("По индексу %d имя не найдено\n", thIdx)
		}
	}

	for _, line := range printCache {
		fmt.Printf("%s\n", line)
	}
	if flgLim {
		fmt.Printf("%s\n", "OK")
	} else {
		fmt.Printf("%s\n", "FAIL")
	}
}
