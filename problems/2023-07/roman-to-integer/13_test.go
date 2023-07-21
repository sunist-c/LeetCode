package roman_to_integer

import (
	"leetcode/cases"
	"testing"
)

func TestRomanToInteger(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[string, int]{
		{
			Input:      "III",
			WantOutput: 3,
		},
		{
			Input:      "IV",
			WantOutput: 4,
		},
		{
			Input:      "IX",
			WantOutput: 9,
		},
		{
			Input:      "LVIII",
			WantOutput: 58,
		},
		{
			Input:      "MCMXCIV",
			WantOutput: 1994,
		},
		{
			Input:      "MMMXLV",
			WantOutput: 3045,
		},
		{
			Input:      "MMMCMXCIX",
			WantOutput: 3999,
		},
		{
			Input:      "MMMCDXCIX",
			WantOutput: 3499,
		},
	}

	t.Run(cases.NewCommonTestCases("RomanToInteger", romanToInt, testCases...).JudgeFunction())
}
