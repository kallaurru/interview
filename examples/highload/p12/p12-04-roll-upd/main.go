package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// parseValue возвращает (число, это_процент).
func parseValue(s string) (int64, bool) {
	s = strings.TrimSpace(s)
	if strings.HasSuffix(s, "%") {
		v, _ := strconv.ParseInt(strings.TrimSuffix(s, "%"), 10, 64)
		return v, true
	}
	v, _ := strconv.ParseInt(s, 10, 64)
	return v, false
}

func main() {
	params := make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		params[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	replicas, _ := strconv.ParseInt(params["replicas"], 10, 64)

	surgeVal, surgePct := parseValue(params["maxSurge"])
	unavailVal, unavailPct := parseValue(params["maxUnavailable"])

	var s, u int64
	if surgePct {
		// maxSurge округляется вверх
		s = (replicas*surgeVal + 99) / 100
	} else {
		s = surgeVal
	}
	if unavailPct {
		// maxUnavailable округляется вниз
		u = replicas * unavailVal / 100
	} else {
		u = unavailVal
	}

	if s == 0 && u == 0 {
		fmt.Println("invalid")
		return
	}

	maxPods := replicas + s
	minAvailable := replicas - u
	if minAvailable < 0 {
		minAvailable = 0
	}

	step := s + u
	batches := (replicas + step - 1) / step // ceil

	fmt.Printf("surge=%d unavailable=%d maxPods=%d minAvailable=%d batches=%d\n",
		s, u, maxPods, minAvailable, batches)
}
