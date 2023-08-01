package find_the_index_of_the_first_occurrence_in_a_string

import (
	"leetcode/cases"
	"testing"
)

type findTheIndexOfTheFirstOccurrenceInAStringInput struct {
	haystack string
	needle   string
}

func TestFindTheIndexOfTheFirstOccurrenceInAString(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[findTheIndexOfTheFirstOccurrenceInAStringInput, int]{
		{
			Input:      findTheIndexOfTheFirstOccurrenceInAStringInput{"hello", "ll"},
			WantOutput: 2,
		},
		{
			Input:      findTheIndexOfTheFirstOccurrenceInAStringInput{"aaaaa", "bba"},
			WantOutput: -1,
		},
		{
			Input:      findTheIndexOfTheFirstOccurrenceInAStringInput{"", ""},
			WantOutput: 0,
		},
	}

	originFunction := func(input findTheIndexOfTheFirstOccurrenceInAStringInput) (output int) {
		return strStr(input.haystack, input.needle)
	}

	cases.NewCommonTestCases("TestFindTheIndexOfTheFirstOccurrenceInAString", originFunction, testCases...).Run(t)
}
