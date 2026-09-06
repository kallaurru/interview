package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var (
		k, q int
	)

	_, err := fmt.Fscan(r, &k, &q)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(1)
	}
	head := make(map[string][]string, k)
	idxItems := make(map[string][]string, k)
	cache := make(map[string]string, k)

	for i := 0; i < k; i++ {
		var idxName, idxPartsLine string

		_, err = fmt.Fscan(r, &idxName, &idxPartsLine)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(1)
		}

		cache[idxPartsLine] = idxName
		idxParts := strings.Split(idxPartsLine, ",")
		if len(idxParts) < 1 {
			continue
		}
		idxlist, ok := head[idxParts[0]]
		if !ok {
			idxlist = make([]string, 0, 2)
		}
		idxlist = append(idxlist, idxName)
		head[idxParts[0]] = idxlist
		idxItems[idxName] = idxParts
	}

	for i := 0; i < q; i++ {
		var qPartsLine string

		_, err = fmt.Fscan(r, &qPartsLine)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(1)
		}
		items := strings.Split(qPartsLine, ",")
		if len(items) == 0 {
			continue
		}
		idx, ok := cache[qPartsLine]
		if ok {
			// префикс совпал полностью
			fmt.Printf("%s %d\n", idx, len(idxItems[idx]))
			continue
		}
		// первое поле префикса
		idxs, ok := head[items[0]]
		if !ok {
			// индексы по указанному префиксу не найдены
			fmt.Printf("%s\n", "FULLSCAN")
			continue
		}
		fit(idxs, idxItems, items)
	}
}

func fit(idxlist []string, idxItems map[string][]string, qFields []string) {
	maxPrefix := 0
	idxNameEnd := ""

	for _, idxName := range idxlist {
		prefixParts := idxItems[idxName]
		counter := 0
		// один проход будет всегда, потому что мы нашли список индексов по первому полю
		for idx := 0; idx < len(qFields); idx++ {
			if prefixParts[idx] != qFields[idx] {
				break
			}
			counter++
			if idx+1 == len(prefixParts) {
				break // закончились поля префикса
			}
		}
		if counter > maxPrefix {
			maxPrefix = counter
			idxNameEnd = idxName
		}
	}
	if maxPrefix == 0 {
		fmt.Printf("%s\n", "FULLSCAN")
		return

	}
	fmt.Printf("%s %d\n", idxNameEnd, maxPrefix)
}
