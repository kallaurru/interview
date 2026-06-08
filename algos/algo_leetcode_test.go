package algos

import (
	"fmt"
	"github.com/kallaurru/interview/algos/leetcode/algo_1"
	"github.com/kallaurru/interview/algos/leetcode/algo_101"
	"github.com/kallaurru/interview/algos/leetcode/algo_125"
	"github.com/kallaurru/interview/algos/leetcode/algo_1446"
	"github.com/kallaurru/interview/algos/leetcode/algo_19"
	"github.com/kallaurru/interview/algos/leetcode/algo_2"
	"github.com/kallaurru/interview/algos/leetcode/algo_20"
	"github.com/kallaurru/interview/algos/leetcode/algo_206"
	"github.com/kallaurru/interview/algos/leetcode/algo_21"
	"github.com/kallaurru/interview/algos/leetcode/algo_22"
	"github.com/kallaurru/interview/algos/leetcode/algo_228"
	"github.com/kallaurru/interview/algos/leetcode/algo_232"
	"github.com/kallaurru/interview/algos/leetcode/algo_234"
	"github.com/kallaurru/interview/algos/leetcode/algo_26"
	"github.com/kallaurru/interview/algos/leetcode/algo_268"
	"github.com/kallaurru/interview/algos/leetcode/algo_283"
	"github.com/kallaurru/interview/algos/leetcode/algo_3"
	"github.com/kallaurru/interview/algos/leetcode/algo_350"
	"github.com/kallaurru/interview/algos/leetcode/algo_387"
	"github.com/kallaurru/interview/algos/leetcode/algo_392"
	"github.com/kallaurru/interview/algos/leetcode/algo_415"
	"github.com/kallaurru/interview/algos/leetcode/algo_5"
	"github.com/kallaurru/interview/algos/leetcode/algo_557"
	"github.com/kallaurru/interview/algos/leetcode/algo_704"
	"github.com/kallaurru/interview/algos/leetcode/algo_771"
	"github.com/kallaurru/interview/algos/leetcode/algo_88"
	"github.com/kallaurru/interview/algos/leetcode/algo_938"
	"github.com/kallaurru/interview/algos/leetcode/algo_977"
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

func TestLeetCodeProblem232(t *testing.T) {
	q := algo_232.New()
	q.Push(1)
	q.Push(2)
	expected := 1
	actual := q.Peek()
	assert.Equal(t, expected, actual, "values not equal")
	actual = q.Pop()
	assert.Equal(t, expected, actual, "values not equal")
	assert.Equal(t, false, q.Empty(), "queue is empty")
}

func TestLeetCodeProblem88(t *testing.T) {
	idx := 0
	for {
		data := GetFixturesAlgo88(idx)
		algo_88.MergeArraysAlgo88(data.Left, data.Right, data.M, data.N)
		assert.Equal(t, data.Expected, data.Left, "arrays not equal")

		idx++
		if !data.Ok {
			break
		}
	}
}

func TestLeetCodeProblem21(t *testing.T) {
	idx := 0
	for {
		data := GetFixtureAlgo21(idx)
		l1 := algo_21.BuildList(data.Left)
		l2 := algo_21.BuildList(data.Right)
		actual := algo_21.MergeTwoListAlgo21(l1, l2)
		i := 0
		for actual != nil {
			assert.Equal(t, data.Expected[i], actual.Val, "values not equal")
			i++
			actual = actual.Next
		}
		if !data.Ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem771(t *testing.T) {
	idx := 0
	for {
		data, expected := GetFixturesAlgo771(idx)
		actual := algo_771.NumJewelsInStonesAlgo771(data.Input, data.Output)
		assert.Equal(t, expected, actual)

		if !data.Ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem350(t *testing.T) {
	idx := 0
	for {
		data := GetFixturesAlgo350(idx)
		actual := algo_350.IntersectionTwoArraysAlgo350(data.Left, data.Right)
		assert.Equal(t, data.Expected, actual, "values not equal")

		if !data.Ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem268(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo228(idx)
		actual := algo_268.MissingNumberAlgo268(in)
		assert.Equal(t, expect, actual, "numbers not equal")

		if !ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem392(t *testing.T) {
	idx := 0
	for {
		data, expected := GetFixtureAlgo392(idx)
		actual := algo_392.IsSubsequenceAlgo392(data.Input, data.Output)
		assert.Equal(t, expected, actual, "not equal values", "idx - ", idx)
		if !data.Ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem977(t *testing.T) {
	idx := 0
	for {
		in, expected, ok := GetFixtureAlgo977(idx)
		actual := algo_977.SquaresOfSortedArrays977(in)
		assert.Equal(t, expected, actual, "not equal values", "idx - ", idx)
		if !ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem26(t *testing.T) {
	idx := 0
	for {
		in, expected, ok := GetFixtureAlgo26(idx)
		actual := algo_26.RemoveDuplicateSortedArray26(in)
		assert.Equal(t, expected, actual, "not equal values", "idx - ", idx)
		if !ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem938(t *testing.T) {
	idx := 0
	for {
		data := GetFixtureAlgo938(idx)
		tree := algo_938.BuildTree(data.Raw, data.MV)
		actual := algo_938.RangeSumBST938(tree, data.L, data.H)
		assert.Equal(t, data.Expect, actual, "values is not equal")

		if !data.Ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem557(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo557(idx)
		actual := algo_557.ReverseWordAlgo557(in)
		assert.Equal(t, expect, actual, "values not equal")

		if !ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem415(t *testing.T) {
	idx := 0
	for {
		data := GetFixtureAlgo415(idx)

		actual := algo_415.AddStringsAlgo415(data.L, data.R)
		assert.Equal(t, data.Expect, actual, "values not equal")

		if !data.Ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem234(t *testing.T) {
	idx := 0
	for {
		raw, expect, ok := GetFixtureAlgo234(idx)
		head := algo_234.BuildList(raw)
		actual := algo_234.IsPalindromeAlgo234(head)
		assert.Equal(t, expect, actual, "values not equal")

		if !ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem387(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo387(idx)
		actual := algo_387.FirstUniqueChrAlgo387(in)
		assert.Equal(t, expect, actual, "values not equal")

		if !ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem2(t *testing.T) {
	idx := 0
begin:
	for {
		data := GetFixtureAlgo2(idx)
		l := algo_2.BuildList(data.Left)
		r := algo_2.BuildList(data.Right)
		actual := algo_2.AddTwoNumbers(l, r)
		for idxR, val := range data.Expected {
			if actual == nil {
				t.Error("node is nil")
				break begin
			}
			assert.Equal(t, val, actual.Val, "values not equal. Idx - ", idxR)
			actual = actual.Next
		}
		if !data.Ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem3(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo3(idx)
		actual := algo_3.LongestSubstring(in)
		assert.Equal(t, expect, actual, fmt.Sprintf("values not equal. Idx- %d", idx))

		if !ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem19(t *testing.T) {
	idx := 0
	for {
		data := GetFixtureAlgo19(idx)
		l := algo_19.BuildList(data.Left)
		actual := algo_19.RemoveNthFromEndAlgo19(l, data.N)
		for _, val := range data.Expected {
			if actual == nil && len(data.Expected) > 0 {
				t.Error("list is end")
				break
			}
			if actual == nil {
				break
			}
			assert.Equal(t, val, actual.Val, "values not equal")
			actual = actual.Next
		}

		if !data.Ok {
			break
		}
		idx++
	}
}

func TestLeetCodeProblem20(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo20(idx)
		actual := algo_20.IsValidAlgo20(in)
		assert.Equal(t, expect, actual, "values is not equal")

		if !ok {
			break
		}

		idx++
	}
}

func TestLeetCodeProblem22(t *testing.T) {
	idx := 0
	for {
		in, expect, ok := GetFixtureAlgo22(idx)
		actual := algo_22.GenerateParenthesisAlgo22(in)
		assert.Equal(t, expect, actual, "values is not equal")

		if !ok {
			break
		}

		idx++
	}
}
