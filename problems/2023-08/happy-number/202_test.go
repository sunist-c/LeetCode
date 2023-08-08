package happy_number

import (
	"leetcode/cases"
	"testing"
)

func TestHappyNumber(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[int, bool]{
		{
			Input:      19,
			WantOutput: true,
		},
		{
			Input:      2,
			WantOutput: false,
		},
		{
			Input:      1,
			WantOutput: true,
		},
	}

	cases.NewCommonTestCases("xHappyNumber", isHappy, testCases...).Run(t)
}
