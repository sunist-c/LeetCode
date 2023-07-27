package delete_greatest_value_in_each_row

import (
	"leetcode/cases"
	"testing"
)

func TestDeleteGreatestValueInEachRow(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[][]int, int]{
		{
			Input:      [][]int{{3, 3, 1}, {2, 4, 1}},
			WantOutput: 8,
		},
		{
			Input:      [][]int{{5, 1}, {4, 2}},
			WantOutput: 7,
		},
		{
			Input:      [][]int{{10}},
			WantOutput: 10,
		},
		{
			Input:      [][]int{{1, 2}, {3, 4}},
			WantOutput: 7,
		},
	}

	cases.NewCommonTestCases("TestDeleteGreatestValueInEachRow", deleteGreatestValue, testCases...).Run(t)
}
