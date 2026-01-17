package algos

import (
	"fmt"
	"github.com/kallaurru/interview/algos/yad/midlinepoints"
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

func TestMidLinePoints(t *testing.T) {
	idx := 0
	for {
		fixture, expect, ok := GetPoints(idx)
		actual := midlinepoints.MidLinePoints(fixture)
		assert.Equal(t, expect, actual, fmt.Sprintf("Set %d is not equal", idx))
		if !ok {
			break
		}
		idx++
	}
}
