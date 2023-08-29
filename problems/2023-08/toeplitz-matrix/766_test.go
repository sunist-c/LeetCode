package toeplitz_matrix

import (
	"leetcode/cases"
	"testing"
)

func TestToeplitzMatrix(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[][]int, bool]{
		{
			Input: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			WantOutput: true,
		},
		{
			Input: [][]int{
				{1, 2},
				{2, 2},
			},
			WantOutput: false,
		},
		{
			Input: [][]int{
				{1, 2, 3, 4},
				{5, 1, 9, 3},
				{9, 5, 1, 2},
			},
			WantOutput: false,
		},
		{
			Input: [][]int{
				{1},
			},
			WantOutput: true,
		},
	}

	cases.NewCommonTestCases("ToepLitzMatrix", isToeplitzMatrix, testCases...).Run(t)
}
