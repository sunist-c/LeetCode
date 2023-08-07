package reverse_string

import (
	"leetcode/cases"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[[]byte, []byte]{
		{
			Input:      []byte{'h', 'e', 'l', 'l', 'o'},
			WantOutput: []byte{'o', 'l', 'l', 'e', 'h'},
		},
		{
			Input:      []byte{'H', 'a', 'n', 'n', 'a', 'h'},
			WantOutput: []byte{'h', 'a', 'n', 'n', 'a', 'H'},
		},
	}

	judgeFunction := func(input []byte, output, wantOutput []byte) bool {
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

	originalSolution := func(input []byte) []byte {
		copiedInput := make([]byte, len(input))
		copy(copiedInput, input)
		reverseString(copiedInput)
		return copiedInput
	}

	cases.NewSpecialJudgeTestCases("ReverseString", originalSolution, judgeFunction, testCases...).Run(t)
}
