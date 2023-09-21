package find_greatest_common_divisor_of_array

import (
	"leetcode/cases"
	"testing"
)

func TestFindGreatestCommonDivisorOfArray(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]int, int]{
		{
			Input: []int{
				2, 5, 6, 9, 10,
			},
			WantOutput: 2,
		},
		{
			Input: []int{
				7, 5, 6, 8, 3,
			},
			WantOutput: 1,
		},
		{
			Input: []int{
				3, 3,
			},
			WantOutput: 3,
		},
	}

	cases.NewCommonTestCases("FindGreatestCommonDivisorOfArray", findGCD, testCases...).Run(t)
}
