package algos

type FixtureNearElements struct {
	Data     []int
	Expected []int
	Idx      int
	K        int
	Ok       bool
}

func GetFixtureNearElements(idx int) *FixtureNearElements {
	// 3th param == false if not found data for idx
	switch idx {
	case 0:
		return &FixtureNearElements{
			Data:     []int{2, 3, 5, 7, 11},
			Expected: []int{7, 5},
			Idx:      3,
			K:        2,
			Ok:       true,
		}
	case 1:
		return &FixtureNearElements{
			Data:     []int{4, 12, 15, 15, 24},
			Expected: []int{12, 15, 15},
			Idx:      1,
			K:        3,
			Ok:       true,
		}
	case 2:
		return &FixtureNearElements{
			Data:     []int{4, 12, 15, 15, 24},
			Expected: []int{4, 12},
			Idx:      0,
			K:        2,
			Ok:       true,
		}
	case 3:
		return &FixtureNearElements{
			Data:     []int{4, 12, 15, 15, 24},
			Expected: []int{24, 15},
			Idx:      4,
			K:        2,
			Ok:       true,
		}
	case 4:
		return &FixtureNearElements{
			Data:     []int{4, 12, 15, 15, 24},
			Expected: []int{},
			Idx:      2,
			K:        0,
			Ok:       true,
		}
	default:
		return &FixtureNearElements{
			Data:     []int{2, 3, 5, 7, 11},
			Expected: []int{5, 7},
			Idx:      2,
			K:        2,
			Ok:       false,
		}
	}
}
