package algos

import (
	"fmt"
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
