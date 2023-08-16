package find_the_losers_of_the_circular_game

import (
	"leetcode/cases"
	"testing"
)

type findTheLosersOfTheCircularGameInput struct {
	n int
	k int
}

func TestFindTheLosersOfTheCircularGame(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[findTheLosersOfTheCircularGameInput, []int]{
		{
			Input: findTheLosersOfTheCircularGameInput{
				n: 5,
				k: 2,
			},
			WantOutput: []int{4, 5},
		},
		{
			Input: findTheLosersOfTheCircularGameInput{
				n: 4,
				k: 4,
			},
			WantOutput: []int{2, 3, 4},
		},
		{
			Input: findTheLosersOfTheCircularGameInput{
				n: 2,
				k: 1,
			},
			WantOutput: []int{},
		},
	}

	originFunction := func(input findTheLosersOfTheCircularGameInput) []int {
		return circularGameLosers(input.n, input.k)
	}

	judgeFunction := func(input findTheLosersOfTheCircularGameInput, output, wantOutput []int) bool {
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

	cases.NewSpecialJudgeTestCases("FindTheLosersOfTheCircularGame", originFunction, judgeFunction, testCases...).Run(t)
}
