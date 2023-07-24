package prime_number_of_set_bits_in_binary_representation

import (
	"leetcode/cases"
	"testing"
)

type primeNumberOfSetBitsInBinaryRepresentationInput struct {
	left  int
	right int
}

func TestPrimeNumberOfSetBitsInBinaryRepresentation(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[primeNumberOfSetBitsInBinaryRepresentationInput, int]{
		{
			Input:      primeNumberOfSetBitsInBinaryRepresentationInput{6, 10},
			WantOutput: 4,
		},
		{
			Input:      primeNumberOfSetBitsInBinaryRepresentationInput{10, 15},
			WantOutput: 5,
		},
	}

	originFunction := func(input primeNumberOfSetBitsInBinaryRepresentationInput) (output int) {
		return countPrimeSetBits(input.left, input.right)
	}

	cases.NewCommonTestCases("PrimeNumberOfSetBitsInBinaryRepresentation", originFunction, testCases...).Run(t)
}
