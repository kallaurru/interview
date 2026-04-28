package algos

import (
	"fmt"
	"github.com/kallaurru/interview/algos/leetcode/algo_125"
	"github.com/kallaurru/interview/algos/leetcode/algo_228"
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
