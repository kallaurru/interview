package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type About struct {
	Name string
	Flat int
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	const limit = 80.0

	_, err := fmt.Fscan(r, &n)
	if err != nil {
		os.Exit(1)
	}
	if n <= 0 {
		os.Exit(11)
	}
	data := make([]About, 0, n)
	sum := 0
	for i := 0; i < n; i++ {
		var name string
		var flat int

		_, err = fmt.Fscan(r, &name, &flat)
		if err != nil {
			fmt.Printf("Ошибка чтения строки #%d: %v\n", i+1, err)
			os.Exit(2)
		}
		about := About{
			Name: name,
			Flat: flat,
		}
		data = append(data, about)
		sum += flat
	}
	sort.Slice(data, func(i, j int) bool {
		// Сначала сравниваем Flat по убыванию
		if data[i].Flat != data[j].Flat {
			return data[i].Flat > data[j].Flat // true — i идёт раньше j
		}
		// При равных Flat — по Name лексикографически (по возрастанию)
		return data[i].Name < data[j].Name
	})
	accum := float64(0)
	for _, item := range data {
		part := float64(item.Flat) / float64(sum) * 100
		fmt.Printf("%s %.1f\n", item.Name, part)
		accum += part
		if accum >= limit {
			break
		}
	}
}
