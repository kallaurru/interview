package algo_771

func NumJewelsInStonesAlgo771(jewels string, stones string) int {
	jewelsStor := make(map[byte]int, len(jewels))
	for i := 0; i < len(jewels); i++ {
		jewelsStor[jewels[i]] = 0
	}

	for i := 0; i < len(stones); i++ {
		_, ok := jewelsStor[stones[i]]
		if ok {
			jewelsStor[stones[i]] += 1
		}
	}

	result := 0
	for _, val := range jewelsStor {
		result += val
	}

	return result
}
