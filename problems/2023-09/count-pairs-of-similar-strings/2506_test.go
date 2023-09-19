package count_pairs_of_similar_strings

import (
	"leetcode/cases"
	"testing"
)

func TestCountPairsOfSimilarStrings(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]string, int]{
		{
			Input:      []string{"abcd", "abcd", "abcd", "abcd"},
			WantOutput: 6,
		},
		{
			Input:      []string{"aba", "aabb", "abcd", "bac", "aabc"},
			WantOutput: 2,
		},
		{
			Input:      []string{"nba", "cba", "dba"},
			WantOutput: 0,
		},
	}

	cases.NewCommonTestCases("CountPairsOfSimilarStrings", similarPairs, testCases...).Run(t)
}
