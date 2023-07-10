package three_sum

import (
	"leetcode/cases"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[[]int, [][]int]{
		{
			Input:      []int{0, 0, 0, 0},
			WantOutput: [][]int{{0, 0, 0}},
		},
		{
			Input:      []int{-1, 0, 1, 2, -1, -4},
			WantOutput: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	judgeFunction := func(input []int, output [][]int, wantOutput [][]int) bool {
		if len(output) != len(wantOutput) {
			return false
		}

		for i, numbers := range wantOutput {
			for j, number := range numbers {
				if output[i][j] != number {
					return false
				}
			}
		}

		return true
	}

	t.Run(cases.NewSpecialJudgeTestCases("ThreeSum", threeSum, judgeFunction, testCases...).JudgeFunction())
}
