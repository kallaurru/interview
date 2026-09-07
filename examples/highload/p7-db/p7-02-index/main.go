package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	idxList := make([]string, 0, k)
	idxItems := make(map[string][]string, k)

	for i := 0; i < k; i++ {
		var idxName, idxPartsLine string

		_, err = fmt.Fscan(r, &idxName, &idxPartsLine)
		if err != nil {
			fmt.Printf("Error - %v", err)
			os.Exit(1)
		}
		idxList = append(idxList, idxName)
		idxItems[idxName] = strings.Split(idxPartsLine, ",")
	}
	sort.Slice(idxList, func(i, j int) bool { return idxList[i] < idxList[j] })

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
		fit(idxList, idxItems, items)
	}
}

func fit(idxlist []string, idxItems map[string][]string, qFields []string) {
	maxPrefix := 0
	idxNameEnd := ""

	for _, idxName := range idxlist {
		items := idxItems[idxName]
		counter := 0
		for _, qItem := range items {
			flg := false
			for idx := 0; idx < len(qFields); idx++ {
				if qItem == qFields[idx] {
					flg = true
					counter++
					break
				}
			}
			if !flg {
				break // прекращаем поиск
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
