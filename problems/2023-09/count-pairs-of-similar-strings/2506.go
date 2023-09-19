package count_pairs_of_similar_strings

func similarPairs(words []string) (result int) {
	counter := map[int]int{}
	for _, word := range words {
		mask := 0
		for _, s := range word {
			mask = mask | (1 << (s - 'a'))
		}
		result += counter[mask]
		counter[mask]++
	}

	return result
}
