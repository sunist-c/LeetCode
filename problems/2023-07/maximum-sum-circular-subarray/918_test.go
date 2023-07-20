package maximum_sum_circular_subarray

import (
	"leetcode/cases"
	"testing"
)

func TestMaximumSumCircularSubarray(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]int, int]{
		{
			Input:      []int{1, -2, 3, -2},
			WantOutput: 3,
		},
		{
			Input:      []int{5, -3, 5},
			WantOutput: 10,
		},
		{
			Input:      []int{3, -1, 2, -1},
			WantOutput: 4,
		},
		{
			Input:      []int{-3, -2, -3},
			WantOutput: -2,
		},
	}

	t.Run(cases.NewCommonTestCases("MaximumSumCircularSubarray", maxSubarraySumCircular, testCases...).JudgeFunction())
}
