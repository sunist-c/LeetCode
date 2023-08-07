package climbing_stairs

import (
	"leetcode/cases"
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[int, int]{
		{
			Input:      2,
			WantOutput: 2,
		},
		{
			Input:      3,
			WantOutput: 3,
		},
		{
			Input:      10,
			WantOutput: 89,
		},
		{
			Input:      45,
			WantOutput: 1836311903,
		},
	}

	cases.NewCommonTestCases("ClimbingStairs", climbStairs, testCases...).Run(t)
}
