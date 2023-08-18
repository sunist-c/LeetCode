package check_permutation_icci

func CheckPermutation(s1 string, s2 string) bool {
	runesOfS1, runesOfS2 := []rune(s1), []rune(s2)
	if len(runesOfS1) != len(runesOfS2) {
		return false
	}

	mapOfS1, mapOfS2 := map[rune]int{}, map[rune]int{}
	for i := 0; i < len(runesOfS1); i++ {
		mapOfS1[runesOfS1[i]]++
		mapOfS2[runesOfS2[i]]++
	}

	for k, v := range mapOfS1 {
		if mapOfS2[k] != v {
			return false
		}
	}

	return true
}
