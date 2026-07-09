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

func maintest() {
	sortList := make([]uint32, 0, constDefStorSize)
	orderingDummy(&sortList)
	fmt.Println(len(sortList))

}
func mainold() {
	idx := 0
	readerIn := bufio.NewReaderSize(getTestReader(idx), constDefBufSize)
	// читаем количество протоколов
	part, err := readerIn.ReadSlice('\n')
	if err != nil {
		os.Exit(1)
	}

	count := parseInt(part)
	for i := 0; i < count; i++ {
		part, err = readerIn.ReadSlice('\n')
		if err != nil {
			os.Exit(1)
		}
		countLines := parseInt(part)
		for l := 0; l < countLines; l++ {
			part, err = readerIn.ReadSlice('\n')
			if err == nil || (err != nil && err == io.EOF) {
				unit, result, ok := converter(part)
				if ok {
					//
					fmt.Printf("UnitID - %d Result - %d\n", unit, result)
				}
			}
			if err != nil {
				os.Exit(1)
			}
		}
	}
	os.Exit(0)
}

func main() {
	idx := 1
	readerIn := bufio.NewReaderSize(getTestReader(idx), constDefBufSize)
	wr := bufio.NewWriterSize(os.Stdout, constDefBufSize)
	defer func() {
		wr.Flush()
	}()
	orderList := make([]uint32, 0, constDefStorSize)
	score := make(map[uint32]uint32, constDefStorSize)
	protoStream, _, ok := reader(readerIn)
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

func reader(r *bufio.Reader) (<-chan Proto, int, bool) {
	chOut := make(chan Proto, 64)
	// читаем количество протоколов
	part, err := r.ReadSlice('\n')
	if err != nil {
		close(chOut)
		return chOut, 0, false
	}
	count := parseInt(part)

	go func(r *bufio.Reader, count int) {
		defer close(chOut)

		for i := 0; i < count; i++ {
			part, err = r.ReadSlice('\n')
			if err != nil {
				return
			}

			countLines := parseInt(part)
			for l := 0; l < countLines; l++ {
				part, err = r.ReadSlice('\n')
				if err == nil || (err != nil && err == io.EOF) {
					unit, result, ok := converter(part)
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

	return chOut, 0, true
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

// todo dummy
func sort(orderList []uint32, values []uint32) []uint32 {
	orderList = append(orderList, values...)

	return orderList
}

func orderingDummy(in *[]uint32) {
	stor := [][]uint32{{1, 3, 5}, {2, 3, 4, 6}, {6, 14, 17}}
	exists := make(map[uint32]struct{}, 10)
	local := *in
	part := make([]uint32, 0, 12)

	for _, series := range stor {
		for _, item := range series {
			_, ok := exists[item]
			if ok {
				continue
			}
			part = append(part, item)
			exists[item] = struct{}{}
		}
		if len(part) > 0 {
			local = sort(local, part)
		}
		part = make([]uint32, 0, 12)
	}
	*in = local
}
