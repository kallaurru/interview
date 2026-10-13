package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

type LogEntry struct {
	Ts        int64  `json:"ts"`
	Level     string `json:"level"`
	Service   string `json:"service"`
	RequestID string `json:"request_id"`
	Msg       string `json:"msg"`
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	// request_id -> list of entries
	logs := make(map[string][]LogEntry)

	for sc.Scan() {
		line := sc.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		var entry LogEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			continue
		}
		if entry.RequestID == "" {
			continue
		}

		logs[entry.RequestID] = append(logs[entry.RequestID], entry)
	}

	type result struct {
		reqID string
		route string
	}
	var results []result

	for reqID, entries := range logs {
		// Check if any entry has level ERROR
		hasError := false
		for _, e := range entries {
			if e.Level == "ERROR" {
				hasError = true
				break
			}
		}
		if !hasError {
			continue
		}

		// Sort by ts ascending
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Ts < entries[j].Ts
		})

		// Build service sequence with consecutive duplicates collapsed
		var services []string
		var prev string
		for i, e := range entries {
			if i == 0 || e.Service != prev {
				services = append(services, e.Service)
				prev = e.Service
			}
		}

		results = append(results, result{
			reqID: reqID,
			route: strings.Join(services, "->"),
		})
	}

	// Sort by request_id lexicographically
	sort.Slice(results, func(i, j int) bool {
		return results[i].reqID < results[j].reqID
	})

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for _, r := range results {
		fmt.Fprintf(out, "%s %s\n", r.reqID, r.route)
	}
}
