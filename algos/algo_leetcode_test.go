package algos

import (
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
