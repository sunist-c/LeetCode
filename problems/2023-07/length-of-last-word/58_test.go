package length_of_last_word

import (
	"leetcode/cases"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[string, int]{
		{
			Input:      "Hello World",
			WantOutput: 5,
		},
		{
			Input:      "   fly me   to   the moon  ",
			WantOutput: 4,
		},
		{
			Input:      "luffy is still joyboy",
			WantOutput: 6,
		},
		{
			Input:      "a",
			WantOutput: 1,
		},
		{
			Input:      "                ",
			WantOutput: 0,
		},
	}

	cases.NewCommonTestCases("TestLengthOfLastWord", lengthOfLastWord, testCases...).Run(t)
}
