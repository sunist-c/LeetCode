package k_items_with_the_maximum_sum

import (
	"leetcode/cases"
	"testing"
)

type kItemsWithTheMaximumSumInput struct {
	numsOne    int
	numsZero   int
	numsNegOne int
	k          int
}

func TestKItemsWithTheMaximumSum(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[kItemsWithTheMaximumSumInput, int]{
		{
			Input:      kItemsWithTheMaximumSumInput{3, 2, 0, 2},
			WantOutput: 2,
		},
		{
			Input:      kItemsWithTheMaximumSumInput{3, 2, 0, 4},
			WantOutput: 3,
		},
		{
			Input:      kItemsWithTheMaximumSumInput{0, 0, 3, 3},
			WantOutput: -3,
		},
		{
			Input:      kItemsWithTheMaximumSumInput{3, 0, 3, 4},
			WantOutput: 2,
		},
	}

	originalFunction := func(input kItemsWithTheMaximumSumInput) (output int) {
		return kItemsWithMaximumSum(input.numsOne, input.numsZero, input.numsNegOne, input.k)
	}

	t.Run(cases.NewCommonTestCases("KItemsWithTheMaximumSum", originalFunction, testCases...).JudgeFunction())
}
