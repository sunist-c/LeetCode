package search_insert_position

import (
	"leetcode/cases"
	"testing"
)

type searchInsertPositionInput struct {
	nums   []int
	target int
}

func TestSearchInsertPosition(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[searchInsertPositionInput, int]{
		{
			Input:      searchInsertPositionInput{[]int{1, 3, 5, 6}, 5},
			WantOutput: 2,
		},
		{
			Input:      searchInsertPositionInput{[]int{1, 3, 5, 6}, 2},
			WantOutput: 1,
		},
		{
			Input:      searchInsertPositionInput{[]int{1, 3, 5, 6}, 7},
			WantOutput: 4,
		},
		{
			Input:      searchInsertPositionInput{[]int{1, 3, 5, 6}, 0},
			WantOutput: 0,
		},
		{
			Input:      searchInsertPositionInput{[]int{1}, 0},
			WantOutput: 0,
		},
	}

	originFunction := func(input searchInsertPositionInput) (output int) {
		return searchInsert(input.nums, input.target)
	}

	cases.NewCommonTestCases("TestSearchInsertPosition", originFunction, testCases...).Run(t)
}
