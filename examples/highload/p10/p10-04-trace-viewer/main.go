package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Span struct {
	id       string
	parentID string
	name     string
	start    int64
	end      int64
	children []*Span
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	spans := make(map[string]*Span)
	var root *Span

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 5 {
			continue
		}
		start, err1 := strconv.ParseInt(parts[3], 10, 64)
		end, err2 := strconv.ParseInt(parts[4], 10, 64)
		if err1 != nil || err2 != nil {
			continue
		}
		s := &Span{
			id:       parts[0],
			parentID: parts[1],
			name:     parts[2],
			start:    start,
			end:      end,
		}
		spans[s.id] = s
	}

	// Build tree
	for _, s := range spans {
		if s.parentID == "-" {
			root = s
			continue
		}
		if parent, ok := spans[s.parentID]; ok {
			parent.children = append(parent.children, s)
		}
	}

	// Sort children by start
	for _, s := range spans {
		sort.Slice(s.children, func(i, j int) bool {
			return s.children[i].start < s.children[j].start
		})
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var dfs func(s *Span, depth int)
	dfs = func(s *Span, depth int) {
		duration := s.end - s.start
		var childSum int64
		for _, c := range s.children {
			childSum += c.end - c.start
		}
		self := duration - childSum

		indent := strings.Repeat("  ", depth)
		fmt.Fprintf(out, "%s%s %d %d\n", indent, s.name, duration, self)

		for _, c := range s.children {
			dfs(c, depth+1)
		}
	}

	if root != nil {
		dfs(root, 0)
	}
}
