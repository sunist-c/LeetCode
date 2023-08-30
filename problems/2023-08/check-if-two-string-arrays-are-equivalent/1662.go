package check_if_two_string_arrays_are_equivalent

import "strings"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	stringBuilderOfWord1, stringBuilderOfWord2 := strings.Builder{}, strings.Builder{}
	for _, word := range word1 {
		stringBuilderOfWord1.WriteString(word)
	}

	for _, word := range word2 {
		stringBuilderOfWord2.WriteString(word)
	}

	return stringBuilderOfWord1.String() == stringBuilderOfWord2.String()
}
