package alternating_digit_sum

import (
	"leetcode/cases"
	"testing"
)

func TestAlternateDigitSum(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[int, int]{
		{
			Input:      521,
			WantOutput: 4,
		},
		{
			Input:      111,
			WantOutput: 1,
		},
		{
			Input:      0,
			WantOutput: 0,
		},
		{
			Input:      886996,
			WantOutput: 0,
		},
		{
			Input:      92323830,
			WantOutput: 6,
		},
	}

	t.Run(cases.NewCommonTestCases("AlternateDigitSum", alternateDigitSum, testCases...).JudgeFunction())
}
