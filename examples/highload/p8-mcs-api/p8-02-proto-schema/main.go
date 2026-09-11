package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	schemaOld = "OLD"
	schemaNew = "NEW"
)

type SchemaItem struct {
	Number int
	Name   string
	Type   string
}

func BuildSchemaItem(line string) (SchemaItem, error) {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return SchemaItem{}, errors.New("bad line")
	}
	numb, err := strconv.Atoi(fields[1])
	if err != nil {
		return SchemaItem{}, err
	}
	return SchemaItem{Number: numb, Name: fields[0], Type: fields[2]}, nil
}

func main() {
	oldSchema, newSchema, reserved := uploadData()
	flg := true
	for _, item := range oldSchema {
		newItem, ok := findItem(item.Number, newSchema)
		if !ok {
			if !findReservedItem(item.Number, reserved) {
				flg = false
				fmt.Printf("REMOVED %d\n", item.Number)
				continue
			}
			continue
		}
		if newItem.Type == item.Type {
			continue
		}

		fmt.Printf("TYPE %d\n", item.Number)
		flg = false
	}
	if flg {
		fmt.Printf("%s\n", "OK")
	}
}

func findItem(n int, stor []SchemaItem) (SchemaItem, bool) {
	for i := 0; i < len(stor); i++ {
		if stor[i].Number == n {
			return stor[i], true
		}
	}
	return SchemaItem{}, false
}

func findReservedItem(n int, stor []int) bool {
	for i := 0; i < len(stor); i++ {
		if stor[i] == n {
			return true
		}
	}
	return false
}

func uploadData() ([]SchemaItem, []SchemaItem, []int) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	next := func() string {
		if !sc.Scan() {
			return ""
		}
		return sc.Text()
	}
	line := next()
	f := strings.Fields(line)

	n, err := strconv.Atoi(f[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}
	oldSchema := make([]SchemaItem, 0, n)
	register := make(map[int]SchemaItem, n)
	for i := 0; i < n; i++ {
		schItem, err := BuildSchemaItem(next())
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(21)
		}
		oldSchema = append(oldSchema, schItem)
		register[schItem.Number] = schItem
	}
	f = strings.Fields(next())
	n, err = strconv.Atoi(f[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(11)
	}
	sort.Slice(oldSchema, func(i, j int) bool {
		return oldSchema[i].Number < oldSchema[j].Number
	})
	newSchema := make([]SchemaItem, 0, n)
	// new schema
	for i := 0; i < n; i++ {
		schItem, err := BuildSchemaItem(next())
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(21)
		}
		newSchema = append(newSchema, schItem)
	}
	sort.Slice(newSchema, func(i, j int) bool {
		return newSchema[i].Number < newSchema[j].Number
	})

	// reserved
	var reserved []int
	fields := strings.Fields(next())
	if len(fields) == 2 {
		// нет зарезервированных
		reserved = nil
	} else {
		for i := 2; i < len(fields); i++ {
			val, err := strconv.Atoi(fields[i])
			if err != nil {
				fmt.Printf("%v\n", err)
				os.Exit(31)
			}
			reserved = append(reserved, val)
		}
	}

	return oldSchema, newSchema, reserved
}
