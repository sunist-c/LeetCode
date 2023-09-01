package number_of_ways_to_buy_pens_and_pencils

import (
	"leetcode/cases"
	"testing"
)

type numbersOfWaysToBuyPensAndPencilsInput struct {
	total int
	cost1 int
	cost2 int
}

func TestNumbersOfWaysToBuyPensAndPencils(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[numbersOfWaysToBuyPensAndPencilsInput, int64]{
		{
			Input:      numbersOfWaysToBuyPensAndPencilsInput{total: 4, cost1: 2, cost2: 3},
			WantOutput: 4,
		},
		{
			Input:      numbersOfWaysToBuyPensAndPencilsInput{total: 6, cost1: 4, cost2: 5},
			WantOutput: 3,
		},
		{
			Input:      numbersOfWaysToBuyPensAndPencilsInput{total: 1, cost1: 1, cost2: 1},
			WantOutput: 3,
		},
	}

	originFunction := func(input numbersOfWaysToBuyPensAndPencilsInput) int64 {
		return waysToBuyPensPencils(input.total, input.cost1, input.cost2)
	}

	cases.NewCommonTestCases("NumberOfWaysToBuyPensAndPencils", originFunction, testCases...).Run(t)
}
