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

func GetFixtureAlgo3(idx int) (string, int, bool) {
	switch idx {
	default:
		return "abcabcbb", 3, false
	case 0:
		return "bbbbb", 1, true
	case 1:
		return "pwwkew", 3, true
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

func GetFixtureAlgo2(idx int) PaarFixture {
	switch idx {
	default:
		return PaarFixture{
			Ok:       false,
			Left:     []int{2, 4, 3},
			Right:    []int{5, 6, 4},
			Expected: []int{7, 0, 8},
		}
	case 0:
		return PaarFixture{
			Ok:       true,
			Left:     []int{0},
			Right:    []int{0},
			Expected: []int{0},
		}
	case 1:
		return PaarFixture{
			Ok:       true,
			Left:     []int{9, 9, 9, 9, 9, 9, 9},
			Right:    []int{9, 9, 9, 9},
			Expected: []int{8, 9, 9, 9, 0, 0, 0, 1},
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

func GetFixtureAlgo392(idx int) (StringFixture, bool) {
	switch idx {
	default:
		return StringFixture{
			Input:  "abc",
			Output: "ahbgdc",
			Ok:     false,
		}, true
	case 0:
		return StringFixture{
			Input:  "axc",
			Output: "ahbgdc",
			Ok:     true,
		}, false
	case 1:
		return StringFixture{
			Input:  "ace",
			Output: "abcde",
			Ok:     true,
		}, true
	case 2:
		return StringFixture{
			Input:  "aec",
			Output: "abcde",
			Ok:     true,
		}, false
	}
}

func GetFixtureAlgo977(idx int) ([]int, []int, bool) {
	switch idx {
	default:
		return []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}, false
	case 0:
		return []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}, true
	}
}

func GetFixtureAlgo26(idx int) ([]int, int, bool) {
	switch idx {
	default:
		return []int{1, 1, 2}, 2, false
	case 0:
		return []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, true
	}
}

func GetFixtureAlgo938(idx int) TreeFixture {
	mv := -1
	switch idx {
	default:
		return TreeFixture{
			Ok:     false,
			MV:     mv,
			Raw:    []int{10, 5, 15, 3, 7, mv, 18},
			L:      7,
			H:      15,
			Expect: 32,
		}
	case 0:
		return TreeFixture{
			Ok:     true,
			MV:     mv,
			Raw:    []int{10, 5, 15, 3, 7, 13, 18, 1, mv, 6, mv, mv, mv, mv, mv},
			L:      6,
			H:      10,
			Expect: 23,
		}
	}
}

func GetFixtureAlgo557(idx int) (string, string, bool) {
	switch idx {
	default:
		return "Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc", false
	case 0:
		return "Mr Ding", "rM gniD", true
	}
}

func GetFixtureAlgo415(idx int) StringFixtureAdv {
	switch idx {
	default:
		return StringFixtureAdv{
			L:      "11",
			R:      "123",
			Expect: "134",
			Ok:     false,
		}
	case 0:
		return StringFixtureAdv{
			L:      "456",
			R:      "77",
			Expect: "533",
			Ok:     true,
		}
	case 1:
		return StringFixtureAdv{
			L:      "0",
			R:      "0",
			Expect: "0",
			Ok:     true,
		}
	}
}

func GetFixtureAlgo234(idx int) ([]int, bool, bool) {
	switch idx {
	default:
		return []int{1, 2, 2, 1}, true, false
	case 0:
		return []int{1, 2}, false, true
	case 1:
		return []int{1, 2, 3, 2, 1}, true, true
	}
}

func GetFixtureAlgo387(idx int) (string, int, bool) {
	switch idx {
	default:
		return "leetcode", 0, false
	case 0:
		return "loveleetcode", 2, true
	case 1:
		return "aabb", -1, true
	}
}

func GetFixtureAlgo19(idx int) PaarFixture {
	switch idx {
	default:
		return PaarFixture{
			Ok:       false,
			Left:     []int{1, 2, 3, 4, 5},
			Expected: []int{1, 2, 3, 5},
			N:        2,
		}
	case 0:
		return PaarFixture{
			Ok:       true,
			Left:     []int{1},
			Expected: []int{},
			N:        1,
		}
	case 1:
		return PaarFixture{
			Ok:       true,
			Left:     []int{1, 2},
			Expected: []int{1},
			N:        1,
		}
	}
}

func GetFixtureAlgo20(idx int) (string, bool, bool) {
	switch idx {
	default:
		return "()", true, false
	case 0:
		return "()[]{}", true, true
	case 1:
		return "(]", false, true
	case 2:
		return "([])", true, true
	case 3:
		return "([)]", false, true
	}
}

func GetFixtureAlgo22(idx int) (int, []string, bool) {
	switch idx {
	default:
		return 1, []string{"()"}, false
	case 0:
		return 3, []string{"()()()", "((()))", "(()())", "(())()", "()(())"}, true
	case 1:
		return 4, []string{"()()()()", "(((())))", "(()()())", "(()())()", "()(()())", "(())()()", "()()(())"}, true
	case 2:
		return 2, []string{"()()", "(())"}, true
	}
}

func GetFixtureAlgo33(idx int) FixtureConfig {
	switch idx {
	default:
		return FixtureConfig{
			Data: []int{4, 5, 6, 7, 0, 1, 2},
			K:    0,
			Idx:  4,
			Ok:   false,
		}
	case 0:
		return FixtureConfig{
			Data: []int{4, 5, 6, 7, 0, 1, 2},
			K:    3,
			Idx:  -1,
			Ok:   true,
		}
	case 1:
		return FixtureConfig{
			Data: []int{1},
			K:    0,
			Idx:  -1,
			Ok:   true,
		}
	}
}

func GetFixtureAlgo49(idx int) ([]string, [][]string, bool) {
	switch idx {
	default:
		return []string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}, false
	case 0:
		return []string{""}, [][]string{{""}}, true
	case 1:
		return []string{"a"}, [][]string{{"a"}}, true
	}
}

func GetFixtureAlgo56(idx int) ([][]int, [][]int, bool) {
	switch idx {
	default:
		return [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}, false
	case 0:
		return [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}, true
	case 1:
		return [][]int{{4, 7}, {1, 4}}, [][]int{{1, 7}}, true
	}
}

func GetFixtureAlgo71(idx int) (string, string, bool) {
	switch idx {
	default:
		return "/home/", "/home", false
	case 0:
		return "/home//foo/", "/home/foo", true
	case 1:
		return "/home/user/Documents/../Pictures", "/home/user/Pictures", true
	case 2:
		return "/../", "/", true
	case 3:
		return "/.../a/../b/c/../d/./", "/.../b/d", true
	}
}
