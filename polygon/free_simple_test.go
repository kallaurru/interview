package polygon

import (
	"fmt"
	"github.com/kallaurru/interview/algos/leetcode/algo_232"
	"github.com/kallaurru/interview/polygon/entities/carrier"
	"github.com/kallaurru/interview/polygon/generates"
	"github.com/kallaurru/interview/polygon/sort"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestMyStack(t *testing.T) {
	stack := algo_232.NewStack()
	in := []int{2, 3, 6, 8, 9, 12}
	for _, val := range in {
		stack.Push(val)
	}
	for i := len(in) - 1; i >= 0; i-- {
		expected := in[i]
		actual := stack.Peek()
		assert.Equal(t, expected, actual, "Values not equal")
		if !stack.IsEmpty() {
			actual = stack.Pop()
			assert.Equal(t, expected, actual, "Values not equal")
		}
	}
	assert.Equal(t, true, stack.IsEmpty(), "stack is not empty")
}

func TestTransformIdxVectorToMatrix(t *testing.T) {
	tr := func(idx, cols int) (int, int) {
		r := idx/cols + 1
		rem := idx % cols
		c := cols - (cols - rem) + 1

		return r, c
	}

	cols := 4
	input := []int{12, 15, 22, 17, 200, 141, 72, 90, 99}

	for idx, val := range input {
		r, c := tr(idx, cols)
		fmt.Printf("Idx - %d, Val - %d, Row - %d, Col - %d\n", idx, val, r, c)
	}
}

func TestBytesManipulate(t *testing.T) {
	var s uint32 = 0x01

	a := s
	fmt.Printf("Step 1. a - %d, s - %d\n", a, s)

	a ^= s
	fmt.Printf("Step 2. a - %d, after ^= byte op\n", a)

	a ^= s
	fmt.Printf("Step 3. a - %d, after ^= byte op\n", a)
}

func TestByteDiffs(t *testing.T) {
	var line = 0
	var result uint32 = 0
	for line = 0; line < 16; line++ {
		result = 1 << line
	}

	log.Println("Result is - ", result)
}

func TestMyStackForAlgo(t *testing.T) {
	stack := algo_232.NewStack()
	for i := 8; i < 16; i++ {
		stack.Push(i)
	}
	fmt.Printf("Size of stack: %d\n", stack.Size())
	for !stack.IsEmpty() {
		fmt.Printf("Val : %d | Size : %d\n", stack.Pop(), stack.Size())
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{1493, 443, 560, 356, 281, 161, 49, 380, 146, 22, 567, 933, 20, 487, 849, 362, 56, 253, 438, 470, 3, 98, 200, 341,
		986, 2, 1004, 5, 236, 150, 71, 19, 279, 1650, 159, 33, 167, 153, 763, 238, 523}
	fmt.Println("Before:", arr)
	sorted := sort.QuickSort(arr)
	fmt.Println("After:", sorted)
}

func TestCarrierSync(t *testing.T) {
	headers := []string{"ticker", "value", "source", "moment", "idx"}

	lines := [][]string{
		{"BTC/USDT", "65000.00", "Bitget", "10:00", "0"},
		{"ETH/USDT", "1498.00", "Bybit", "11:15", "1"},
		{"MNT/USDT", "0.46431", "Bybit", "14:15", "2"}}

	carri := carrier.NewCarr(4, headers)
	for _, part := range lines {
		carri.AddLine(part)
	}

	expectedLines := len(lines) + 1 // заголовки в наличии
	assert.Equal(t, expectedLines, carri.Lines(), "count of lines not equal")
	assert.True(t, carri.Len() > 0, "len of carrier is zero")

	fmt.Printf("Len -%d, Lines - %d\n", carri.Len(), carri.Lines())
	data := carri.Parse()
	for i := 0; i < carri.Lines(); i++ {
		item, ok := data[i]
		if !ok {
			continue
		}
		fmt.Printf("Begin line idx %d ----------- \n", i)
		for idx, val := range item {
			if idx == carri.KeyF() {
				fmt.Printf("Idx F - %d | Is Index - Yes | Val - %s\n", idx, string(val))
			} else {
				fmt.Printf("Idx F - %d | Is Index - No | Val - %s\n", idx, string(val))
			}
		}
		fmt.Printf("End of Line ----------- \n\n")
	}
}

func TestCarrierUnmarshall(t *testing.T) {
	headers := []string{"ticker", "value", "source", "moment", "idx"}

	lines := [][]string{
		{"BTC/USDT", "65000.00", "Bitget", "10:00", "0"},
		{"ETH/USDT", "1498.00", "Bybit", "11:15", "1"},
		{"MNT/USDT", "0.46431", "Bybit", "14:15", "2"}}

	carri := carrier.NewCarr(4, headers)
	for _, part := range lines {
		carri.AddLine(part)
	}

	expectedLines := len(lines) + 1 // заголовки в наличии
	assert.Equal(t, expectedLines, carri.Lines(), "count of lines not equal")
	assert.True(t, carri.Len() > 0, "len of carrier is zero")

	fmt.Printf("Len -%d, Lines - %d\n", carri.Len(), carri.Lines())

	data := carrier.Unmarshall(carri)
	lineIdx := 0
	for lineAsFields := range data {
		fmt.Printf("Begin line idx %d ----------- \n", lineIdx)
		for idx, val := range lineAsFields {
			if idx == carri.KeyF() {
				fmt.Printf("Idx F - %d | Is Index - Yes | Val - %s\n", idx, string(val))
			} else {
				fmt.Printf("Idx F - %d | Is Index - No | Val - %s\n", idx, string(val))
			}
		}
		fmt.Printf("End of Line ----------- \n\n")
		lineIdx++

	}
}

func TestReqIDGen(t *testing.T) {
	const uuid uint16 = 0x0808

	reqID, ok := generates.GenerateReqID(uuid)
	if !ok {
		os.Exit(2)
	}
	str, ok := generates.ParseNiceReqID(reqID)
	if !ok {
		os.Exit(3)
	}
	fmt.Printf("%s\n", str)

}
