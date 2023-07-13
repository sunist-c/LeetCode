package minimum_falling_path_sum

import (
	"leetcode/cases"
	"testing"
)

func TestMinimumFallingPathSum(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[][]int, int]{
		{
			Input:      [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}},
			WantOutput: 13,
		},
		{
			Input:      [][]int{{-19, 57}, {-40, -5}},
			WantOutput: -59,
		},
	}

	t.Run(cases.NewCommonTestCases("MinimumFallingPathSum", minFallingPathSum, testCases...).JudgeFunction())
}
