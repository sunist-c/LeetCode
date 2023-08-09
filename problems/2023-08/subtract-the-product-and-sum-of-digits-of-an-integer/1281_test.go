package subtract_the_product_and_sum_of_digits_of_an_integer

import (
	"leetcode/cases"
	"testing"
)

func TestSubtractTheProductAndSumOfDigitsOfAnInteger(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[int, int]{
		{
			Input:      234,
			WantOutput: 15,
		},
		{
			Input:      4421,
			WantOutput: 21,
		},
		{
			Input:      1,
			WantOutput: 0,
		},
	}

	cases.NewCommonTestCases("SubtractTheProductAndSumOfDigitsOfAnInteger", subtractProductAndSum, testCases...).Run(t)
}
