package polygon

import (
	"fmt"
	"github.com/kallaurru/interview/algos/leetcode/algo_232"
	"log"
	"testing"
)

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
