package algos

import "image"

func GetFixtureNearElements(idx int) *FixtureConfig {
	// 3th param == false if not found data for idx
	switch idx {
	case 0:
		return &FixtureConfig{
			Data:     []int{2, 3, 5, 7, 11},
			Expected: []int{7, 5},
			Idx:      3,
			K:        2,
			Ok:       true,
		}
	case 1:
		return &FixtureConfig{
			Data:     []int{4, 12, 15, 15, 24},
			Expected: []int{12, 15, 15},
			Idx:      1,
			K:        3,
			Ok:       true,
		}
	case 2:
		return &FixtureConfig{
			Data:     []int{4, 12, 15, 15, 24},
			Expected: []int{4, 12},
			Idx:      0,
			K:        2,
			Ok:       true,
		}
	case 3:
		return &FixtureConfig{
			Data:     []int{4, 12, 15, 15, 24},
			Expected: []int{24, 15},
			Idx:      4,
			K:        2,
			Ok:       true,
		}
	case 4:
		return &FixtureConfig{
			Data:     []int{4, 12, 15, 15, 24},
			Expected: []int{},
			Idx:      2,
			K:        0,
			Ok:       true,
		}
	default:
		return &FixtureConfig{
			Data:     []int{2, 3, 5, 7, 11},
			Expected: []int{5, 7},
			Idx:      2,
			K:        2,
			Ok:       false,
		}
	}
}

func GetPoints(idx int) ([]image.Point, bool, bool) {
	switch idx {
	case 0:
		return []image.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 1}, {X: 4, Y: 1}, {X: 4, Y: 0}}, true, true
	case 1:
		return []image.Point{{X: 0, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 1}, {X: 4, Y: 1}, {X: 4, Y: 0}}, false, true
	case 2:
		return []image.Point{{X: 0, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 1}, {X: 4, Y: 0}}, false, true
	case 3:
		return []image.Point{{X: 0, Y: 0}, {X: 10, Y: 0}}, true, true
	case 4:
		return []image.Point{{X: 0, Y: 0}, {X: 11, Y: 1}}, false, true
	case 5:
		return []image.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 4, Y: 0}}, false, true
	default:
		return []image.Point{{X: 0, Y: 0}}, false, false
	}
}

func GetKSumFixtures(idx int) *FixtureConfig {
	switch idx {
	default:
		return &FixtureConfig{
			Data:     []int{11, 3, 7, 3, 2, 9, 3},
			Expected: []int{2, 4},
			K:        12,
			Ok:       false,
		}
	}
}
