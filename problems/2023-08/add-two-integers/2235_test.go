package add_two_integers

import (
	"leetcode/cases"
	"testing"
)

type addTwoIntegersInput struct {
	num1 int
	num2 int
}

func TestAddTwoIntegers(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[addTwoIntegersInput, int]{
		{
			Input: addTwoIntegersInput{
				num1: 1,
				num2: 2,
			},
			WantOutput: 3,
		},
		{
			Input: addTwoIntegersInput{
				num1: 0,
				num2: 0,
			},
			WantOutput: 0,
		},
		{
			Input: addTwoIntegersInput{
				num1: 1,
				num2: -1,
			},
			WantOutput: 0,
		},
		{
			Input: addTwoIntegersInput{
				num1: -1,
				num2: -1,
			},
			WantOutput: -2,
		},
	}

	originFunction := func(input addTwoIntegersInput) int {
		return sum(input.num1, input.num2)
	}

	cases.NewCommonTestCases("AddTwoIntegers", originFunction, testCases...).Run(t)
}
