package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024*1024), 1024*1024)

	seen := make(map[string]bool)
	likes := make(map[string]map[string]bool)

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}
		parts := strings.Fields(line)
		if len(parts) < 4 {
			continue
		}
		eventID := parts[0]
		userID := parts[1]
		postID := parts[2]
		action := parts[3]

		if seen[eventID] {
			continue
		}
		seen[eventID] = true

		if action == "LIKE" {
			if likes[postID] == nil {
				likes[postID] = make(map[string]bool)
			}
			likes[postID][userID] = true
		} else if action == "UNLIKE" {
			if likes[postID] != nil {
				delete(likes[postID], userID)
				if len(likes[postID]) == 0 {
					delete(likes, postID)
				}
			}
		}
	}

	var posts []string
	for postID, users := range likes {
		if len(users) > 0 {
			posts = append(posts, postID)
		}
	}

	sort.Strings(posts)

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for _, postID := range posts {
		count := len(likes[postID])
		fmt.Fprintf(out, "%s=%d\n", postID, count)
	}
}
