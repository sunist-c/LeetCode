package sum_in_a_matrix

import (
	"leetcode/cases"
	"testing"
)

func TestSumInAMatrix(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[][]int, int]{
		{
			Input:      [][]int{{1, 2, 3}, {3, 2, 1}, {3, 1, 5}},
			WantOutput: 9,
		},
		{
			Input:      [][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}},
			WantOutput: 15,
		},
		{
			Input:      [][]int{{1}},
			WantOutput: 1,
		},
	}

	t.Run(cases.NewCommonTestCases("SumInAMatrix", matrixSum, testCases...).JudgeFunction())
}
