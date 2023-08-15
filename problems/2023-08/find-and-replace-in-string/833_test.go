package find_and_replace_in_string

import (
	"leetcode/cases"
	"testing"
)

type findAndReplaceInStringInput struct {
	s       string
	indices []int
	sources []string
	targets []string
}

func TestFindAndReplaceInString(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[findAndReplaceInStringInput, string]{
		{
			Input: findAndReplaceInStringInput{
				s:       "abcd",
				indices: []int{0, 2},
				sources: []string{"a", "cd"},
				targets: []string{"eee", "ffff"},
			},
			WantOutput: "eeebffff",
		},
		{
			Input: findAndReplaceInStringInput{
				s:       "abcd",
				indices: []int{0, 2},
				sources: []string{"ab", "ec"},
				targets: []string{"eee", "ffff"},
			},
			WantOutput: "eeecd",
		},
	}

	originFunction := func(input findAndReplaceInStringInput) string {
		return findReplaceString(input.s, input.indices, input.sources, input.targets)
	}

	cases.NewCommonTestCases("FindAndReplaceInString", originFunction, testCases...).Run(t)
}
