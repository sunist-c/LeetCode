package card_flipping_game

import (
	"leetcode/cases"
	"testing"
)

type cardFlippingGameInput struct {
	fronts []int
	backs  []int
}

func TestCardFlippingGame(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[cardFlippingGameInput, int]{
		{
			Input:      cardFlippingGameInput{fronts: []int{1, 2, 4, 4, 7}, backs: []int{1, 3, 4, 1, 3}},
			WantOutput: 2,
		},
		{
			Input:      cardFlippingGameInput{fronts: []int{1, 1}, backs: []int{1, 2}},
			WantOutput: 2,
		},
		{
			Input:      cardFlippingGameInput{fronts: []int{1}, backs: []int{1}},
			WantOutput: 0,
		},
	}

	originFunction := func(input cardFlippingGameInput) (output int) {
		return flipgame(input.fronts, input.backs)
	}

	cases.NewCommonTestCases("CardFlippingGame", originFunction, testCases...).Run(t)
}
