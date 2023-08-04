package add_binary

import (
	"leetcode/cases"
	"testing"
)

type addBinaryInput struct {
	a string
	b string
}

func TestAddBinary(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[addBinaryInput, string]{
		{
			Input:      addBinaryInput{a: "11", b: "1"},
			WantOutput: "100",
		},
		{
			Input:      addBinaryInput{a: "1010", b: "1011"},
			WantOutput: "10101",
		},
		{
			Input:      addBinaryInput{a: "0", b: "0"},
			WantOutput: "0",
		},
	}

	originFunction := func(input addBinaryInput) (output string) {
		return addBinary(input.a, input.b)
	}

	cases.NewCommonTestCases("AddBinary", originFunction, testCases...).Run(t)
}
