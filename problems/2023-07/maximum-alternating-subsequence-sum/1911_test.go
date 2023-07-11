package maximum_alternating_subsequence_sum

import (
	"leetcode/cases"
	"testing"
)

func TestMaximumAlternatingSubsequenceSum(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]int, int64]{
		{
			Input:      []int{4, 2, 5, 3},
			WantOutput: 7,
		},
		{
			Input:      []int{5, 6, 7, 8},
			WantOutput: 8,
		},
		{
			Input:      []int{6, 2, 1, 2, 4, 5},
			WantOutput: 10,
		},
		{
			Input:      []int{4, 2, 1, 2, 4, 5},
			WantOutput: 8,
		},
		{
			Input:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			WantOutput: 9,
		},
		{
			Input:      []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			WantOutput: 9,
		},
	}

	t.Run(cases.NewCommonTestCases("MaximumAlternatingSubsequenceSum", maxAlternatingSum, testCases...).JudgeFunction())
}
