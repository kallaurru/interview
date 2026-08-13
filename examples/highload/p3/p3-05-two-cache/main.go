package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stat struct {
	DBReads int
	Stale   int
	Hits    int
}

var db map[string]int

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		fmt.Printf("Error - %v", err)
		os.Exit(2)
	}

	cacheA := make(map[string]int, (n+1)/2)
	cacheB := make(map[string]int, (n+1)/2)
	db = make(map[string]int, (n+1)/2)
	statA := Stat{}
	statB := Stat{}

	for i := 0; i < n; i++ {
		// разбираем линию на запчасти
		var op, key string

		_, err = fmt.Fscan(r, &op, &key)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(3)
		}
		switch op {
		default:
			continue
		case "SET":
			setVal(key)
			delete(cacheB, key)
		case "GET":
			statA = strategyA(cacheA, key, statA)
			statB = strategyB(cacheB, key, statB)
		}
	}

	fmt.Printf("A: hits=%d stale=%d db_reads=%d\n", statA.Hits, statA.Stale, statA.DBReads)
	fmt.Printf("B: hits=%d db_reads=%d\n", statB.Hits, statB.DBReads)
}

func strategyA(cache map[string]int, key string, stat Stat) Stat {
	version, ok := cache[key]
	dbVer := getVal(key)
	if !ok {
		stat.DBReads++
		cache[key] = dbVer
		return stat
	}
	stat.Hits++
	if version < dbVer {
		stat.Stale++
	}

	return stat
}

func strategyB(cache map[string]int, key string, stat Stat) Stat {
	_, ok := cache[key]
	if ok {
		stat.Hits++
		return stat
	}
	stat.DBReads++
	dbVersion := getVal(key)
	cache[key] = dbVersion

	return stat
}

func setVal(key string) int {
	version, ok := db[key]
	if !ok {
		db[key] = 0
		return 0
	}

	db[key] += 1
	return version + 1
}

func getVal(key string) int {
	version, ok := db[key]
	if !ok {
		v := setVal(key)
		return v
	}

	return version
}
