package longest_common_prefix

import (
	"leetcode/cases"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]string, string]{
		{
			Input:      []string{"flower", "flow", "flight"},
			WantOutput: "fl",
		},
		{
			Input:      []string{"dog", "racecar", "car"},
			WantOutput: "",
		},
		{
			Input:      []string{"ab", "a"},
			WantOutput: "a",
		},
	}

	cases.NewCommonTestCases("LongestCommonPrefix", longestCommonPrefix, testCases...).Run(t)
}
