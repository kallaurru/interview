package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseCPU переводит CPU в милликоры. Возвращает (значение, задано ли поле).
func parseCPU(s string) (int64, bool) {
	if s == "-" {
		return 0, false
	}
	if strings.HasSuffix(s, "m") {
		v, err := strconv.ParseInt(s[:len(s)-1], 10, 64)
		if err != nil {
			return 0, false
		}
		return v, true
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, false
	}
	return v * 1000, true
}

// parseMem переводит память в Mi (1 Gi = 1024 Mi).
func parseMem(s string) (int64, bool) {
	if s == "-" {
		return 0, false
	}
	if strings.HasSuffix(s, "Gi") {
		v, err := strconv.ParseInt(s[:len(s)-2], 10, 64)
		if err != nil {
			return 0, false
		}
		return v * 1024, true
	}
	if strings.HasSuffix(s, "Mi") {
		v, err := strconv.ParseInt(s[:len(s)-2], 10, 64)
		if err != nil {
			return 0, false
		}
		return v, true
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, false
	}
	return v, true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)

	if !scanner.Scan() {
		return
	}
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		return
	}

	anyResourceSet := false // хотя бы одно поле задано у любого контейнера
	allGuaranteed := true   // все контейнеры удовлетворяют Guaranteed

	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		fields := strings.Fields(scanner.Text())
		if len(fields) != 5 {
			continue
		}

		cpuReq, cpuReqSet := parseCPU(fields[1])
		cpuLim, cpuLimSet := parseCPU(fields[2])
		memReq, memReqSet := parseMem(fields[3])
		memLim, memLimSet := parseMem(fields[4])

		if cpuReqSet || cpuLimSet || memReqSet || memLimSet {
			anyResourceSet = true
		}

		// Гарантированно: должны быть заданы limits и по CPU, и по памяти,
		// а requests (после нормализации admission-контроллером) равны limits.
		ok := cpuLimSet && memLimSet

		// Нормализация: если request не задан, а limit задан — request = limit.
		effCpuReq := cpuReq
		if !cpuReqSet {
			effCpuReq = cpuLim
		}
		effMemReq := memReq
		if !memReqSet {
			effMemReq = memLim
		}

		if cpuLimSet && effCpuReq != cpuLim {
			ok = false
		}
		if memLimSet && effMemReq != memLim {
			ok = false
		}

		if !ok {
			allGuaranteed = false
		}
	}

	switch {
	case !anyResourceSet:
		fmt.Println("BestEffort")
	case allGuaranteed:
		fmt.Println("Guaranteed")
	default:
		fmt.Println("Burstable")
	}
}
