package algo_49

import "fmt"

func GroupAnagramsAlgo49(str []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range str {
		count := [26]int{}
		for _, ch := range s {
			count[ch-'a']++
		}
		key := fmt.Sprintf("%v", count)
		groups[key] = append(groups[key], s)
	}

	result := make([][]string, 0, len(groups))
	for _, group := range groups {
		result = append(result, group)
	}

	return result
}
