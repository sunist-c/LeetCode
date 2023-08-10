package element_appearing_more_than_25_in_sorted_array

import (
	"leetcode/cases"
	"testing"
)

func TestElementAppearingMoreThan25InSortedArray(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]int, int]{
		{
			Input:      []int{1, 2, 2, 6, 6, 6, 6, 7, 10},
			WantOutput: 6,
		},
		{
			Input:      []int{},
			WantOutput: 0,
		},
	}

	cases.NewCommonTestCases("ElementAppearingMoreThan25InSortedArray", findSpecialInteger, testCases...).Run(t)
}
