package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Normalizer struct {
	exp *regexp.Regexp
	to  string
}

func New() Normalizer {
	return Normalizer{
		exp: regexp.MustCompile(`\d+`),
		to:  "N",
	}
}

func (n Normalizer) Normalize(q string) string {
	return n.exp.ReplaceAllString(q, n.to)
}

func (n Normalizer) Hash(qRaw string) (string, string) {
	normal := n.Normalize(qRaw)
	hash := md5.Sum([]byte(normal))

	return hex.EncodeToString(hash[:]), normal
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	var (
		line string
	)
	counter := make(map[string]int, 16)
	examples := make(map[string]string, 16)
	forOrder := make([]string, 0, 16)
	n := New()
	for scan.Scan() {
		line = scan.Text()
		parts := strings.Fields(line) // разделяем по пробелам
		if len(parts) < 1 {
			fmt.Printf("Skipping invalid line: %q\n", line)
			continue
		}
		sqlRaw := strings.Join(parts[1:], " ")
		hash, normalized := n.Hash(sqlRaw)
		counter[hash] += 1
		_, ok := examples[hash]
		if !ok {
			forOrder = append(forOrder, hash)
		}
		examples[hash] = normalized
	}
	hasTarget := false
	for _, hash := range forOrder {
		count := counter[hash]
		if count < 10 {
			continue
		}
		example := examples[hash]
		fmt.Printf("%d %s\n", count, example)
		hasTarget = true
	}

	if !hasTarget {
		fmt.Printf("%s\n", "OK")
	}
}
