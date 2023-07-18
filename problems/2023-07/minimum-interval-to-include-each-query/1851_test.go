package minimum_interval_to_include_each_query

import (
	"leetcode/cases"
	"testing"
)

type minimumIntervalToIncludeEachQueryInput struct {
	intervals [][]int
	queries   []int
}

func TestMinimumIntervalToIncludeEachQuery(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[minimumIntervalToIncludeEachQueryInput, []int]{
		{
			Input: minimumIntervalToIncludeEachQueryInput{
				intervals: [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}},
				queries:   []int{2, 3, 4, 5},
			},
			WantOutput: []int{3, 3, 1, 4},
		},
		{
			Input: minimumIntervalToIncludeEachQueryInput{
				intervals: [][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}},
				queries:   []int{2, 19, 5, 22},
			},
			WantOutput: []int{2, -1, 4, 6},
		},
	}

	judgeFunction := func(input minimumIntervalToIncludeEachQueryInput, output []int, wantOutput []int) bool {
		if len(wantOutput) != len(output) {
			return false
		}

		for i, value := range wantOutput {
			if output[i] != value {
				return false
			}
		}

		return true
	}

	originFunction := func(input minimumIntervalToIncludeEachQueryInput) []int {
		return minInterval(input.intervals, input.queries)
	}

	t.Run(cases.NewSpecialJudgeTestCases("MinimumIntervalToIncludeEachQuery", originFunction, judgeFunction, testCases...).JudgeFunction())
}
