package insert_interval

import (
	"leetcode/cases"
	"testing"
)

type insertIntervalInput struct {
	intervals   [][]int
	newInterval []int
}

func TestInsertInterval(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[insertIntervalInput, [][]int]{
		{
			Input: insertIntervalInput{
				intervals: [][]int{
					{1, 3},
				},
				newInterval: []int{2, 5},
			},
			WantOutput: [][]int{
				{1, 5},
			},
		},
		{
			Input: insertIntervalInput{
				intervals: [][]int{
					{1, 3}, {6, 9},
				},
				newInterval: []int{2, 5},
			},
			WantOutput: [][]int{
				{1, 5}, {6, 9},
			},
		},
		{
			Input: insertIntervalInput{
				intervals: [][]int{
					{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
				},
				newInterval: []int{4, 8},
			},
			WantOutput: [][]int{
				{1, 2}, {3, 10}, {12, 16},
			},
		},
		{
			Input: insertIntervalInput{
				intervals: [][]int{
					{1, 5},
				},
				newInterval: []int{5, 7},
			},
			WantOutput: [][]int{
				{1, 7},
			},
		},
	}

	originFunction := func(input insertIntervalInput) (output [][]int) {
		return insert(input.intervals, input.newInterval)
	}

	judgeFunction := func(input insertIntervalInput, output, wantOutput [][]int) bool {
		if len(output) != len(wantOutput) {
			return false
		}

		for i := 0; i < len(output); i++ {
			if output[i][0] != wantOutput[i][0] || output[i][1] != wantOutput[i][1] {
				return false
			}
		}

		return true
	}

	cases.NewSpecialJudgeTestCases("InsertInterval", originFunction, judgeFunction, testCases...).Run(t)
}
