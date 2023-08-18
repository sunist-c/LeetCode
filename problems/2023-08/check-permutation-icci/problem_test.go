package check_permutation_icci

import (
	"leetcode/cases"
	"testing"
)

type checkPermutationInput struct {
	s1 string
	s2 string
}

func TestCheckPermutation(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[checkPermutationInput, bool]{
		{
			Input: checkPermutationInput{
				s1: "abc",
				s2: "bca",
			},
			WantOutput: true,
		},
		{
			Input: checkPermutationInput{
				s1: "abc",
				s2: "bad",
			},
			WantOutput: false,
		},
		{
			Input: checkPermutationInput{
				s1: "a",
				s2: "",
			},
			WantOutput: false,
		},
	}

	originFunction := func(input checkPermutationInput) bool {
		return CheckPermutation(input.s1, input.s2)
	}

	cases.NewCommonTestCases("CheckPermutation", originFunction, testCases...).Run(t)
}
