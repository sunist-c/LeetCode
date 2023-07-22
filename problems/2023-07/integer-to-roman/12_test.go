package integer_to_roman

import (
	"leetcode/cases"
	"testing"
)

func TestIntegerToRoman(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[int, string]{
		{
			Input:      3,
			WantOutput: "III",
		},
		{
			Input:      4,
			WantOutput: "IV",
		},
		{
			Input:      9,
			WantOutput: "IX",
		},
		{
			Input:      58,
			WantOutput: "LVIII",
		},
		{
			Input:      1994,
			WantOutput: "MCMXCIV",
		},
		{
			Input:      3045,
			WantOutput: "MMMXLV",
		},
		{
			Input:      3999,
			WantOutput: "MMMCMXCIX",
		},
	}

	t.Run(cases.NewCommonTestCases("IntegerToRoman", intToRoman, testCases...).JudgeFunction())
}
