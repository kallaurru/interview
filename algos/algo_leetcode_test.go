package algos

import (
	"fmt"
	"github.com/kallaurru/interview/algos/leetcode/algo_1"
	"github.com/kallaurru/interview/algos/leetcode/algo_101"
	"github.com/kallaurru/interview/algos/leetcode/algo_125"
	"github.com/kallaurru/interview/algos/leetcode/algo_1446"
	"github.com/kallaurru/interview/algos/leetcode/algo_206"
	"github.com/kallaurru/interview/algos/leetcode/algo_228"
	"github.com/kallaurru/interview/algos/leetcode/algo_283"
	"github.com/kallaurru/interview/algos/leetcode/algo_5"
	"github.com/kallaurru/interview/algos/leetcode/algo_704"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBinSearchAlgo_704(t *testing.T) {
	i := 0
	for {
		data := GetFixturesBinSearchLeetCode702(i)
		actual := algo_704.BinSearchAlgo704(data.Data, data.K)
		assert.Equal(t, data.Idx, actual)
		i++
		if !data.Ok {
			break
		}
	}
}

func TestLeetCodeProblem5(t *testing.T) {
	i := 8
	for {
		item := GetFixturesAlgo5(i)
		actual := algo_5.MaxPalindromeAlgo5(item.Input)
		assert.Equal(t, item.Output, actual)
		i++
		if !item.Ok {
			break
		}
	}
}

func TestLeetCodeProblem228(t *testing.T) {
	i := 0
	for {
		in, expect, ok := GetFixturesAlgo228(i)
		actual := algo_228.SummaryRanges(in)
		assert.Equal(t, len(expect), len(actual), "lens of arrays not equal")
		assert.Equal(t, expect, actual)
		i++
		if !ok {
			break
		}
	}
	fmt.Printf("Count - %d\n", i)
}

func TestLeetCodeProblem125(t *testing.T) {
	i := 0
	for {
		in, expect, ok := GetFixtureAlgo125(i)
		actual := algo_125.ValidPalindrome(in)
		assert.Equal(t, expect, actual, "not equal. case - ", i)
		i++
		if !ok {
			break
		}
	}
}

func TestLeetCodeProblem283(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo283(idx)
		actual := algo_283.MoveZero(in)
		assert.Equal(t, expect, actual, "arrays not equal")
		idx++
		if !ok {
			break
		}
	}
}

func TestLeetCodeProblem206(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo206(idx)
		actual := algo_206.ReverseListRecursive(in, 0, len(in)-1)
		assert.Equal(t, expect, actual, "arrays not equal. Case recursive")

		in, expect, ok = GetFixtureAlgo206(idx)
		actual = algo_206.ReverseListIteration(in)
		assert.Equal(t, expect, actual, "arrays not equal. Case interation")
		idx++
		if !ok {
			break
		}
	}
}

func TestLeetCodeProblem1(t *testing.T) {
	idx := 0
	for {
		fixt := GetFixturesAlgo1(idx)
		actual := algo_1.TwoSumNear(fixt.Data, fixt.K)
		assert.Equal(t, fixt.Expected, actual, "arrays not equal")
		idx++
		if !fixt.Ok {
			break
		}
	}
}

func TestLeetCodeProblem101(t *testing.T) {
	idx := 0
	mv := algo_101.ConstMissVal
	for {
		in, expect, ok := GetFixtureAlgo101(idx, mv)
		actual := algo_101.SymmetricTreeI(in, mv)
		assert.Equal(t, expect, actual, "tree error I")

		in, expect, ok = GetFixtureAlgo101(idx, mv)
		actual = algo_101.SymmetricTreeR(in, mv)
		assert.Equal(t, expect, actual, "tree error R")
		idx++
		if !ok {
			break
		}
	}
}

func TestLeetCodeProblem1446(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo1446(idx)
		actual := algo_1446.ConsecutiveChar(in)
		assert.Equal(t, expect, actual, "not equal")

		idx++
		if !ok {
			break
		}
	}
}
