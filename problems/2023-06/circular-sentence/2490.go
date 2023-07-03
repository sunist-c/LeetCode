package circular_sentence

import "strings"

func isCircularSentence(sentence string) bool {
	splits := strings.Split(sentence, " ")
	var firstRune, prevRune rune
	for i, split := range splits {
		runeArray := []rune(split)
		lengthOfArray := len(runeArray)
		if lengthOfArray == 0 {
			continue
		}

		if i == 0 {
			firstRune = runeArray[0]
		} else if runeArray[0] != prevRune {
			return false
		}

		if i == len(splits)-1 && firstRune != runeArray[lengthOfArray-1] {
			return false
		}

		prevRune = runeArray[lengthOfArray-1]
	}

	return true
}
