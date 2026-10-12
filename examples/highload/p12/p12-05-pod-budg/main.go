package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	params := make(map[string]string)
	type zone struct {
		name string
		pods int64
	}
	var zones []zone

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1024*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if idx := strings.Index(line, "="); idx >= 0 {
			key := strings.TrimSpace(line[:idx])
			val := strings.TrimSpace(line[idx+1:])
			params[key] = val
		} else {
			fields := strings.Fields(line)
			if len(fields) >= 2 {
				p, _ := strconv.ParseInt(fields[1], 10, 64)
				zones = append(zones, zone{fields[0], p})
			}
		}
	}

	replicas, _ := strconv.ParseInt(params["replicas"], 10, 64)
	healthy, _ := strconv.ParseInt(params["healthy"], 10, 64)

	// Проценты в обеих формах округляются вверх
	ceilPct := func(total, pct int64) int64 {
		return (total*pct + 99) / 100
	}

	var desiredHealthy int64
	if v, ok := params["minAvailable"]; ok {
		if strings.HasSuffix(v, "%") {
			pct, _ := strconv.ParseInt(strings.TrimSuffix(v, "%"), 10, 64)
			desiredHealthy = ceilPct(replicas, pct)
		} else {
			desiredHealthy, _ = strconv.ParseInt(v, 10, 64)
		}
	} else if v, ok := params["maxUnavailable"]; ok {
		var maxU int64
		if strings.HasSuffix(v, "%") {
			pct, _ := strconv.ParseInt(strings.TrimSuffix(v, "%"), 10, 64)
			maxU = ceilPct(replicas, pct)
		} else {
			maxU, _ = strconv.ParseInt(v, 10, 64)
		}
		desiredHealthy = replicas - maxU
	}

	allowed := healthy - desiredHealthy
	if allowed < 0 {
		allowed = 0
	}

	drain := "blocked"
	if allowed >= 1 {
		drain = "ok"
	}

	// Зона с максимумом подов (первая при равенстве)
	var maxName string
	var maxPods int64
	if len(zones) > 0 {
		maxName = zones[0].name
		maxPods = zones[0].pods
		for _, z := range zones[1:] {
			if z.pods > maxPods {
				maxPods = z.pods
				maxName = z.name
			}
		}
	} else {
		maxName = "-"
		maxPods = 0
	}

	zonedrain := "blocked"
	if healthy-maxPods >= desiredHealthy {
		zonedrain = "ok"
	}

	fmt.Printf("desiredHealthy=%d allowed=%d drain=%s maxzone=%s/%d zonedrain=%s\n",
		desiredHealthy, allowed, drain, maxName, maxPods, zonedrain)
}
