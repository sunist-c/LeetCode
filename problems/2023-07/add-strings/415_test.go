package add_strings

import (
	"leetcode/cases"
	"testing"
)

type addStringsInput struct {
	num1 string
	num2 string
}

func TestAddStrings(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[addStringsInput, string]{
		{
			Input: addStringsInput{
				num1: "11",
				num2: "123",
			},
			WantOutput: "134",
		},
		{
			Input: addStringsInput{
				num1: "456",
				num2: "77",
			},
			WantOutput: "533",
		},
		{
			Input: addStringsInput{
				num1: "0",
				num2: "0",
			},
			WantOutput: "0",
		},
		{
			Input: addStringsInput{
				num1: "1",
				num2: "9",
			},
			WantOutput: "10",
		},
		{
			Input: addStringsInput{
				num1: "",
				num2: "",
			},
			WantOutput: "0",
		},
	}

	originFunction := func(input addStringsInput) string {
		return addStrings(input.num1, input.num2)
	}

	t.Run(cases.NewCommonTestCases("AddStrings", originFunction, testCases...).JudgeFunction())
}
