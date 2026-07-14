package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	//"strconv"
	//"sync"
)

type Proto struct {
	UnitID  uint32
	Result  uint32
	ProtoID int
}

const (
	constDefChanSize = 64
	constDefStorSize = 2048
	constDefBufSize  = 1 << 20
)

func main() {
	idx := 1
	readerIn := bufio.NewReaderSize(getTestReader(idx), constDefBufSize)
	wr := bufio.NewWriterSize(os.Stdout, constDefBufSize)
	defer func() {
		wr.Flush()
	}()
	orderList := make([]uint32, 0, constDefStorSize)
	score := make(map[uint32]uint32, constDefStorSize)
	protoStream, ok := reader(readerIn)
	wg := sync.WaitGroup{}
	if !ok {
		return
	}
	wg.Add(3)
	go func(in <-chan Proto, wg *sync.WaitGroup) {
		defer wg.Done()
		chOrder := make(chan Proto, constDefChanSize)
		chCheck := make(chan Proto, constDefChanSize)

		go ordering(&orderList, chOrder, wg)
		go checkin(chCheck, score, wg)
		for proto := range in {
			chOrder <- proto
			chCheck <- proto
		}
		close(chOrder)
		close(chCheck)
	}(protoStream, &wg)
	wg.Wait()
	final(wr, score, orderList)
}

func reader(r *bufio.Reader) (<-chan Proto, bool) {
	chOut := make(chan Proto, 64)
	// читаем количество протоколов
	tmp, err := r.ReadSlice('\n')
	if err != nil {
		close(chOut)
		return chOut, false
	}
	count := parseInt(tmp)
	if count == 0 {
		close(chOut)
		return chOut, false
	}
	go func(r *bufio.Reader, count int) {
		defer close(chOut)

		for i := 0; i < count; i++ {
			tmp, err = r.ReadSlice('\n')
			if err != nil {
				return
			}

			countLines := parseInt(tmp)
			for l := 0; l < countLines; l++ {
				tmp, err = r.ReadSlice('\n')
				if err == nil || (err != nil && err == io.EOF) {
					unit, result, ok := converter(tmp)
					if ok {
						//
						chOut <- Proto{
							ProtoID: i,
							UnitID:  unit,
							Result:  result,
						}
					}
				}
				if err != nil {
					break
				}
			}
		}
	}(r, count)

	return chOut, true
}

func converter(in []byte) (uint32, uint32, bool) {
	parts := bytes.Split(in, []byte{32})
	if len(parts) == 1 {
		return uint32(parseInt(parts[0])), 0, false
	}

	return uint32(parseInt(parts[0])), uint32(parseInt(parts[1])), true
}

func ordering(orderList *[]uint32, in <-chan Proto, wg *sync.WaitGroup) {
	const start = -1
	defer wg.Done()

	score := make(map[uint32]struct{}, 1<<16)
	local := *orderList
	curProtoID := start
	var part []uint32
	for proto := range in {
		if curProtoID == start {
			curProtoID = proto.ProtoID
			part = make([]uint32, 0, constDefStorSize)
		}
		_, ok := score[proto.UnitID]
		if ok {
			continue
		}
		score[proto.UnitID] = struct{}{}
		if proto.ProtoID == curProtoID {
			part = append(part, proto.UnitID)
			continue
		}
		// смена протокола
		if len(part) > 0 {
			local = sort(local, part)
		}
		part = make([]uint32, 0, constDefStorSize)
		part = append(part, proto.UnitID)
		curProtoID = proto.ProtoID
	}
	if len(part) > 0 {
		local = sort(local, part)
	}
	*orderList = local
}

func checkin(in <-chan Proto, score map[uint32]uint32, wg *sync.WaitGroup) {
	defer wg.Done()

	for proto := range in {
		score[proto.UnitID] = proto.Result
	}
}

func final(wr *bufio.Writer, score map[uint32]uint32, orderList []uint32) {
	if len(orderList) == 0 {
		wr.WriteString(fmt.Sprintf("Empty order list\n"))
	}
	for _, unitID := range orderList {
		val, ok := score[unitID]
		if !ok {
			continue
		}
		wr.WriteString(fmt.Sprintf("%d %d\n", unitID, val))
	}
	wr.WriteString("\n")
}

func parseInt(b []byte) int {
	n := 0
	for _, c := range b {
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n
}

func sort(orderList []uint32, values []uint32) []uint32 {
	if len(orderList) == 0 {
		orderList = append(orderList, values...)
		return orderList
	}
	start := findL(orderList, values[0])
	end := findR(orderList, values[len(values)-1])

	return mergeArrays(orderList, values, start, end)
}

// Left: первый индекс, где in[i] >= target
func findL(in []uint32, target uint32) int {
	l, r := 0, len(in)
	for l < r {
		mid := l + (r-l)/2
		if in[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// Right: первый индекс, где in[i] > target
func findR(in []uint32, target uint32) int {
	l, r := 0, len(in)
	for l < r {
		mid := l + (r-l)/2
		if in[mid] <= target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func mergeArrays(l, r []uint32, start, end int) []uint32 {
	mid := mergeSubArrays(l, r, start, end)
	left := l[:start]
	right := l[end:]
	out := make([]uint32, 0, len(left)+len(mid)+len(right))
	out = append(out, left...)
	out = append(out, mid...)
	out = append(out, right...)

	return out
}

func mergeSubArrays(l, r []uint32, start, end int) []uint32 {
	tmp := l[start:end]
	out := make([]uint32, len(tmp)+len(r))
	ptr1, ptr2 := len(tmp)-1, len(r)-1
	for i := len(out) - 1; i >= 0; i-- {
		if ptr1 < 0 && ptr2 >= 0 {
			out[i] = r[ptr2]
			ptr2--
			continue
		}

		if ptr2 < 0 && ptr1 >= 0 {
			out[i] = tmp[ptr1]
			ptr1--
			continue
		}
		if tmp[ptr1] > r[ptr2] {
			out[i] = tmp[ptr1]
			if ptr1 >= 0 {
				ptr1--
			}
			continue
		}
		out[i] = r[ptr2]
		if ptr2 >= 0 {
			ptr2--
		}
	}
	return out
}
