package na_ying_bi

import (
	"leetcode/cases"
	"testing"
)

func TestNaYingBi(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[[]int, int]{
		{
			Input:      []int{4, 2, 1},
			WantOutput: 4,
		},
		{
			Input:      []int{2, 3, 10},
			WantOutput: 8,
		},
		{
			Input:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			WantOutput: 25,
		},
	}

	cases.NewCommonTestCases("NaYingBi", minCount, testCases...).Run(t)
}
