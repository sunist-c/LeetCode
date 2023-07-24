package jewels_and_stones

import (
	"leetcode/cases"
	"testing"
)

type jewelsAndStonesInput struct {
	jewels string
	stones string
}

func TestJewelsAndStones(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[jewelsAndStonesInput, int]{
		{
			Input:      jewelsAndStonesInput{"aA", "aAAbbbb"},
			WantOutput: 3,
		},
		{
			Input:      jewelsAndStonesInput{"z", "ZZ"},
			WantOutput: 0,
		},
		{
			Input:      jewelsAndStonesInput{"", ""},
			WantOutput: 0,
		},
		{
			Input:      jewelsAndStonesInput{"dec", "daaccedebababecbbeacebbdadbceacbacdbbadeaeccbecbac"},
			WantOutput: 26,
		},
	}

	originFunction := func(input jewelsAndStonesInput) (output int) {
		return numJewelsInStones(input.jewels, input.stones)
	}

	t.Run(cases.NewCommonTestCases("JewelsAndStones", originFunction, testCases...).JudgeFunction())
}
