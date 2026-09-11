package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := io.ReadAll(os.Stdin)
	fields := strings.Fields(string(data))
	if len(fields) == 0 {
		return
	}

	rtt, _ := strconv.ParseInt(fields[0], 10, 64)
	tokens := fields[1:]
	pos := 0

	var parse func() (int64, int64)
	parse = func() (int64, int64) {
		tok := tokens[pos]
		pos++
		switch tok {
		case "SEQ":
			k, _ := strconv.Atoi(tokens[pos])
			pos++
			var sum, cnt int64
			for i := 0; i < k; i++ {
				l, c := parse()
				sum += l
				cnt += c
			}
			return sum, cnt
		case "PAR":
			k, _ := strconv.Atoi(tokens[pos])
			pos++
			var mx, cnt int64
			for i := 0; i < k; i++ {
				l, c := parse()
				if l > mx {
					mx = l
				}
				cnt += c
			}
			return mx, cnt
		case "CALL":
			pos++ // пропускаем имя сервиса
			ms, _ := strconv.ParseInt(tokens[pos], 10, 64)
			pos++
			return ms + rtt, 1
		}
		return 0, 0
	}

	lat, cnt := parse()
	fmt.Printf("%d %d\n", lat, cnt)
}
