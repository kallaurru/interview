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

func GetFixtureAlgo283(idx int) ([]int, []int, bool) {
	switch idx {
	default:
		return []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}, false
	case 0:
		return []int{0}, []int{0}, true
	case 1:
		return []int{0, 1, 0, 3, 12, 0}, []int{1, 3, 12, 0, 0, 0}, true
	case 2:
		return []int{1, 2, 3, 12, 0, 20}, []int{1, 2, 3, 12, 20, 0}, true
	}
}

func GetFixtureAlgo206(idx int) ([]int, []int, bool) {
	switch idx {
	default:
		return []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, false
	case 0:
		return []int{1, 2}, []int{2, 1}, true
	case 1:
		return []int{}, []int{}, true
	}
}

func GetFixturesAlgo1(idx int) *FixtureConfig {
	switch idx {
	default:
		return &FixtureConfig{
			Data:     []int{2, 7, 11, 15},
			Expected: []int{0, 1},
			K:        9,
			Ok:       false,
		}
	case 0:
		return &FixtureConfig{
			Data:     []int{3, 2, 4},
			Expected: []int{1, 2},
			K:        6,
			Ok:       true,
		}
	case 1:
		return &FixtureConfig{
			Data:     []int{3, 3},
			Expected: []int{0, 1},
			K:        6,
			Ok:       true,
		}
	}
}

// GetFixtureAlgo101 По условиям задачи макс значение от -100 до 100.
// Значит пропущенные значения обозначаем как
func GetFixtureAlgo101(idx int, missVal int) ([]int, bool, bool) {
	switch idx {
	default:
		return []int{1, 2, 2, 3, 4, 4, 3}, true, false
	case 0:
		return []int{1, 2, 2, missVal, 3, missVal, 3}, false, true
	}
}

func GetFixtureAlgo1446(idx int) (string, int, bool) {
	switch idx {
	default:
		return "leetcode", 2, false
	case 0:
		return "abbcccddddeeeeedcba", 5, true
	}
}

func GetFixturesAlgo88(idx int) PaarFixture {
	switch idx {
	default:
		return PaarFixture{
			Ok:       false,
			Left:     []int{1, 2, 3, 0, 0, 0},
			Right:    []int{2, 5, 6},
			Expected: []int{1, 2, 2, 3, 5, 6},
			M:        3,
			N:        3,
		}
	case 0:
		return PaarFixture{
			Ok:       true,
			Left:     []int{1},
			Right:    []int{},
			Expected: []int{1},
			M:        1,
			N:        0,
		}
	case 1:
		return PaarFixture{
			Ok:       true,
			Left:     []int{0},
			Right:    []int{1},
			Expected: []int{1},
			M:        0,
			N:        1,
		}
	case 2:
		return PaarFixture{
			Ok:       true,
			Left:     []int{1, 2, 5, 0, 0, 0},
			Right:    []int{2, 5, 6},
			Expected: []int{1, 2, 2, 5, 5, 6},
			M:        3,
			N:        3,
		}
	case 3:
		return PaarFixture{
			Ok:       true,
			Left:     []int{4, 5, 8, 0, 0, 0},
			Right:    []int{2, 5, 6},
			Expected: []int{2, 4, 5, 5, 6, 8},
			M:        3,
			N:        3,
		}
	}
}

func GetFixtureAlgo21(idx int) PaarFixture {
	switch idx {
	default:
		return PaarFixture{
			Ok:       false,
			Left:     []int{1, 5, 7, 12},
			Right:    []int{1, 2, 3, 8, 15, 16},
			Expected: []int{1, 1, 2, 3, 5, 7, 8, 12, 15, 16},
		}
	case 0:
		return PaarFixture{
			Ok:       true,
			Left:     []int{1, 2, 4},
			Right:    []int{1, 3, 4},
			Expected: []int{1, 1, 2, 3, 4, 4},
		}
	case 1:
		return PaarFixture{
			Ok:       true,
			Left:     []int{},
			Right:    []int{},
			Expected: []int{},
		}
	case 2:
		return PaarFixture{
			Ok:       true,
			Left:     []int{},
			Right:    []int{0},
			Expected: []int{0},
		}
	}
}

func GetFixturesAlgo771(idx int) (StringFixture, int) {
	switch idx {
	default:
		return StringFixture{
			Input:  "aA",
			Output: "aAAbbbb",
			Ok:     false,
		}, 3
	case 0:
		return StringFixture{
			Input:  "z",
			Output: "ZZ",
			Ok:     true,
		}, 0
	}
}

func GetFixturesAlgo350(idx int) PaarFixture {
	switch idx {
	default:
		return PaarFixture{
			Ok:       false,
			Left:     []int{1, 2, 2, 1},
			Right:    []int{2, 2},
			Expected: []int{2, 2},
		}
	case 0:
		return PaarFixture{
			Ok:       true,
			Left:     []int{4, 9, 5},
			Right:    []int{9, 4, 9, 8, 4},
			Expected: []int{9, 4},
		}
	}
}

func GetFixtureAlgo228(idx int) ([]int, int, bool) {
	switch idx {
	default:
		return []int{3, 0, 1}, 2, false
	case 0:
		return []int{0, 1}, 2, true
	case 1:
		return []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8, true
	}
}
