package water_bottles

import (
	"leetcode/cases"
	"testing"
)

type waterBottlesInput struct {
	numBottles  int
	numExchange int
}

func TestWaterBottles(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[waterBottlesInput, int]{
		{
			Input:      waterBottlesInput{numBottles: 9, numExchange: 3},
			WantOutput: 13,
		},
		{
			Input:      waterBottlesInput{numBottles: 15, numExchange: 4},
			WantOutput: 19,
		},
		{
			Input:      waterBottlesInput{numBottles: 5, numExchange: 5},
			WantOutput: 6,
		},
	}

	originFunction := func(input waterBottlesInput) int {
		return numWaterBottles(input.numBottles, input.numExchange)
	}

	cases.NewCommonTestCases("WaterBottles", originFunction, testCases...).Run(t)
}
