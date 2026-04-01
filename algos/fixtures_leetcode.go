package algos

func GetFixturesBinSearchLeetCode702(idx int) *FixtureConfig {
	switch idx {
	default:
		return &FixtureConfig{
			Data:     []int{-1, 0, 3, 5, 9, 12},
			Expected: nil,
			Idx:      -1, // ожидаемый результат
			K:        2,  // target
			Ok:       false,
		}
	case 0:
		return &FixtureConfig{
			Data:     []int{-1, 0, 3, 5, 9, 12},
			Expected: nil,
			Idx:      4, // ожидаемый результат
			K:        9, // target
			Ok:       true,
		}
	}
}

func GetFixturesAlgo5(idx int) *StringFixture {
	switch idx {
	default:
		return &StringFixture{
			Input:  "babad",
			Output: "bab",
			Ok:     false,
		}
	case 0:
		return &StringFixture{
			Input:  "cbbd",
			Output: "bb",
			Ok:     true,
		}
	case 1:
		return &StringFixture{
			Input:  "babab",
			Output: "babab",
			Ok:     true,
		}
	case 2:
		return &StringFixture{
			Input:  "baddab",
			Output: "baddab",
			Ok:     true,
		}
	case 3:
		return &StringFixture{
			Input:  "baddac",
			Output: "adda",
			Ok:     true,
		}
	}
}
