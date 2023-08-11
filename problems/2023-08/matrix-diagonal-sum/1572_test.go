package matrix_diagonal_sum

import (
	"leetcode/cases"
	"testing"
)

func TestMatrixDiagonalSum(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[][]int, int]{
		{
			Input:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			WantOutput: 25,
		},
		{
			Input:      [][]int{{5}},
			WantOutput: 5,
		},
		{
			Input:      [][]int{},
			WantOutput: 0,
		},
	}

	cases.NewCommonTestCases("MatrixDiagonalSum", diagonalSum, testCases...).Run(t)
}
