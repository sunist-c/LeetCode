package check_if_two_string_arrays_are_equivalent

import (
	"testing"

	"leetcode/cases"
)

type checkIfTwoStringArraysAreEquivalentInput struct {
	word1 []string
	word2 []string
}

func TestCheckIfTwoStringArraysAreEquivalent(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[checkIfTwoStringArraysAreEquivalentInput, bool]{
		{
			Input: checkIfTwoStringArraysAreEquivalentInput{
				word1: []string{"ab", "c"},
				word2: []string{"a", "bc"},
			},
			WantOutput: true,
		},
		{
			Input: checkIfTwoStringArraysAreEquivalentInput{
				word1: []string{"a", "cb"},
				word2: []string{"ab", "c"},
			},
			WantOutput: false,
		},
		{
			Input: checkIfTwoStringArraysAreEquivalentInput{
				word1: []string{"abc", "d", "defg"},
				word2: []string{"abcddefg"},
			},
			WantOutput: true,
		},
		{
			Input: checkIfTwoStringArraysAreEquivalentInput{
				word1: []string{"abc", "d", "defg"},
			},
			WantOutput: false,
		},
	}

	originFunction := func(input checkIfTwoStringArraysAreEquivalentInput) (output bool) {
		return arrayStringsAreEqual(input.word1, input.word2)
	}

	cases.NewCommonTestCases("CheckIfTwoStringArraysAreEquivalent", originFunction, testCases...).Run(t)
}
