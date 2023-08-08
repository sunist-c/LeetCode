package maximum_absolute_sum_of_any_subarray

import (
	"leetcode/cases"
	"testing"
)

func TestMaximumAbsoluteSumOfAnySubarray(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]int, int]{
		{
			Input:      []int{1, -3, 2, 3, -4},
			WantOutput: 5,
		},
		{
			Input:      []int{2, -5, 1, -4, 3, -2},
			WantOutput: 8,
		},
		{
			Input:      []int{1},
			WantOutput: 1,
		},
		{
			Input:      []int{},
			WantOutput: 0,
		},
	}

	cases.NewCommonTestCases("TestMaximumAbsoluteSumOfAnySubarray", maxAbsoluteSum, testCases...).Run(t)
}
