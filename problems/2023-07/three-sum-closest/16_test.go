package three_sum_closest

import (
	"leetcode/cases"
	"testing"
)

type threeSumClosestInput struct {
	numbers []int
	target  int
}

func TestThreeSumClosest(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[threeSumClosestInput, int]{
		{
			Input: threeSumClosestInput{
				numbers: []int{-1, 2, 1, -4},
				target:  1,
			},
			WantOutput: 2,
		},
		{
			Input: threeSumClosestInput{
				numbers: []int{0, 0, 0},
				target:  1,
			},
			WantOutput: 0,
		},
	}

	originFunction := func(input threeSumClosestInput) int {
		return threeSumClosest(input.numbers, input.target)
	}

	t.Run(cases.NewCommonTestCases("ThreeSumClosest", originFunction, testCases...).JudgeFunction())
}
