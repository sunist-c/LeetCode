package move_pieces_to_obtain_a_string

import (
	"leetcode/cases"
	"testing"
)

type movePiecesToObtainAStringInput struct {
	start  string
	target string
}

func TestMovePiecesToObtainAString(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[movePiecesToObtainAStringInput, bool]{
		{
			Input: movePiecesToObtainAStringInput{
				start:  "LR__",
				target: "L__R",
			},
			WantOutput: true,
		},
		{
			Input: movePiecesToObtainAStringInput{
				start:  "LL_RR",
				target: "L__RR",
			},
			WantOutput: false,
		},
		{
			Input: movePiecesToObtainAStringInput{
				start:  "L",
				target: "L",
			},
			WantOutput: true,
		},
	}

	originFunction := func(input movePiecesToObtainAStringInput) bool {
		return canChange(input.start, input.target)
	}

	cases.NewCommonTestCases("MovePiecesToObtainAString", originFunction, testCases...).Run(t)
}
