package merge_strings_alternately

import (
	"leetcode/cases"
	"testing"
)

type mergeStringsAlternatelyInput struct {
	word1 string
	word2 string
}

func TestMergeStringsAlternately(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[mergeStringsAlternatelyInput, string]{
		{
			Input: mergeStringsAlternatelyInput{
				word1: "abc",
				word2: "pqr",
			},
			WantOutput: "apbqcr",
		},
		{
			Input: mergeStringsAlternatelyInput{
				word1: "ab",
				word2: "pqrs",
			},
			WantOutput: "apbqrs",
		},
		{
			Input: mergeStringsAlternatelyInput{
				word1: "abcd",
				word2: "pq",
			},
			WantOutput: "apbqcd",
		},
	}

	originFunction := func(input mergeStringsAlternatelyInput) (output string) {
		return mergeAlternately(input.word1, input.word2)
	}

	cases.NewCommonTestCases("MergeStringsAlternately", originFunction, testCases...).Run(t)
}
