package algos

import (
	"github.com/kallaurru/interview/algos/yad/tnearkfori"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNearElements(t *testing.T) {
	idx := 0
	for {
		fixture := GetFixtureNearElements(idx)
		actual := tnearkfori.NearKElements(fixture.Data, fixture.Idx, fixture.K)
		if len(fixture.Expected) == 0 {
			assert.Equal(t, 0, len(actual), "not equal empty arrays")
		} else {
			assert.Equal(t, fixture.Expected, actual, "not equal input index")
		}
		if !fixture.Ok {
			break
		}
		idx++
	}
}
