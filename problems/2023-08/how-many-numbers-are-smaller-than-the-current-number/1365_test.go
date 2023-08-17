package how_many_numbers_are_smaller_than_the_current_number

import (
	"leetcode/cases"
	"testing"
)

func TestHowManyNumbersAreSmallerThanTheCurrentNumber(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[[]int, []int]{
		{
			Input:      []int{8, 1, 2, 2, 3},
			WantOutput: []int{4, 0, 1, 1, 3},
		},
		{
			Input:      []int{6, 5, 4, 8},
			WantOutput: []int{2, 1, 0, 3},
		},
		{
			Input:      []int{7, 7, 7, 7},
			WantOutput: []int{0, 0, 0, 0},
		},
	}

	judgeFunction := func(input, output, wantOutput []int) bool {
		if len(output) != len(wantOutput) {
			return false
		}

		for i := 0; i < len(output); i++ {
			if output[i] != wantOutput[i] {
				return false
			}
		}

		return true
	}

	cases.NewSpecialJudgeTestCases("HowManyNumbersAreSmallerThanTheCurrentNumber", smallerNumbersThanCurrent, judgeFunction, testCases...).Run(t)
}
