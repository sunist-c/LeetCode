package minimum_operations_to_halve_array_sum

import (
	"leetcode/cases"
	"testing"
)

func TestMinimumOperationsToHalveArraySum(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]int, int]{
		{
			Input:      []int{5, 19, 8, 1},
			WantOutput: 3,
		},
		{
			Input:      []int{3, 8, 20},
			WantOutput: 3,
		},
	}

	cases.NewCommonTestCases("MinimumOperationsToHalveArraySum", halveArray, testCases...).Run(t)
}
