package count_servers_that_communicate

import (
	"leetcode/cases"
	"testing"
)

func TestCountServersThatCommunicate(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[][]int, int]{
		{
			Input:      [][]int{{1, 0}, {0, 1}},
			WantOutput: 0,
		},
		{
			Input:      [][]int{{1, 0}, {1, 1}},
			WantOutput: 3,
		},
		{
			Input:      [][]int{{1, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 1}},
			WantOutput: 4,
		},
	}

	cases.NewCommonTestCases("CountServersThatCommunicate", countServers, testCases...).Run(t)
}
