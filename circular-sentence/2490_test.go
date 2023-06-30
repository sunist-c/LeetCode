package circular_sentence

import (
	"leetcode/cases"
	"testing"
)

func TestCircularSentence(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[string, bool]{
		{
			Input:      "Leetcode is cool",
			WantOutput: false,
		},
		{
			Input:      "leetcode exercises sound delightful",
			WantOutput: true,
		},
		{
			Input:      "eetcode",
			WantOutput: true,
		},
		{
			Input:      "leetcode",
			WantOutput: false,
		},
	}

	t.Run(cases.NewCommonTestCases("CircularSentence", isCircularSentence, testCases...).JudgeFunction())
}
