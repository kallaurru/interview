package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Metric struct {
	name  string
	typ   string
	count int64
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	byName := make(map[string]*Metric)
	var order []string

	for sc.Scan() {
		line := sc.Text()
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}

		// Comment / directive lines
		if strings.HasPrefix(trimmed, "#") {
			fields := strings.Fields(trimmed)
			if len(fields) >= 4 && fields[1] == "TYPE" {
				name := fields[2]
				typ := fields[3]
				if existing, ok := byName[name]; ok {
					existing.typ = typ
				} else {
					m := &Metric{name: name, typ: typ}
					byName[name] = m
					order = append(order, name)
				}
			}
			continue
		}

		// Sample line: <metric>[{labels}] <value>
		lastSpace := strings.LastIndexByte(trimmed, ' ')
		if lastSpace < 0 {
			continue
		}
		metricSpec := trimmed[:lastSpace]

		// Extract metric name (strip labels)
		name := metricSpec
		if idx := strings.IndexByte(metricSpec, '{'); idx >= 0 {
			name = metricSpec[:idx]
		}

		// Prefer exact match, then suffix match
		if m, ok := byName[name]; ok {
			m.count++
			continue
		}
		matched := false
		for _, n := range order {
			if name == n+"_bucket" || name == n+"_sum" || name == n+"_count" {
				byName[n].count++
				matched = true
				break
			}
		}
		_ = matched
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var total int64
	for _, name := range order {
		m := byName[name]
		fmt.Fprintf(out, "%s %s %d\n", m.name, m.typ, m.count)
		total += m.count
	}
	fmt.Fprintf(out, "total %d\n", total)
}
