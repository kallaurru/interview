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
			Input:  "193867",
			Output: "",
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
	case 4:
		return &StringFixture{
			Input:  "bb",
			Output: "bb",
			Ok:     true,
		}
	case 5:
		return &StringFixture{
			Input:  "Abdjscc",
			Output: "cc",
			Ok:     true,
		}
	case 6:
		return &StringFixture{
			Input:  "aba",
			Output: "aba",
			Ok:     true,
		}
	case 7:
		return &StringFixture{
			Input:  "babad",
			Output: "bab",
			Ok:     false,
		}
	}
}

func GetFixturesAlgo228(idx int) ([]int, []string, bool) {
	switch idx {
	default:
		return []int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}, false
	case 0:
		return []int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}, true
	}
}

func GetFixtureAlgo125(idx int) (string, bool, bool) {
	switch idx {
	default:
		return "A man, a plan, a canal: Panama", true, false
	case 0:
		return "race a car", false, true
	case 1:
		return "race aa ecar", true, true
	case 2:
		return "abba", true, true
	}
}
