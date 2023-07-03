package cases

import (
	"fmt"
	"testing"
)

type SpecialJudgeTestCase[InputStructure any, OutputStructure any] struct {
	Input      InputStructure
	WantOutput OutputStructure
}

type SpecialJudgeTestCases[InputStructure any, OutputStructure any] struct {
	name          string
	cases         []SpecialJudgeTestCase[InputStructure, OutputStructure]
	implement     func(InputStructure) OutputStructure
	judgeFunction func(input InputStructure, output, wantOutput OutputStructure) bool
}

func (cases *SpecialJudgeTestCases[InputStructure, OutputStructure]) JudgeFunction() (name string, testingFunction func(t *testing.T)) {
	return fmt.Sprintf("Test%s-SpecialJudge", cases.name),
		func(t *testing.T) {
			for i, testCase := range cases.cases {
				t.Run(fmt.Sprintf("%s-%d", cases.name, i+1), func(t *testing.T) {
					output := cases.implement(testCase.Input)
					if !cases.judgeFunction(testCase.Input, output, testCase.WantOutput) {
						t.Errorf("Failed to Run Test in %s: %d\n"+
							"Input: %#v\n"+
							"Want Output: %#v\n"+
							"Get Output: %#v\n",
							cases.name, i+1, testCase.Input, testCase.WantOutput, output,
						)
					}
				})
			}
		}
}

func NewSpecialJudgeTestCases[InputStructure any, OutputStructure any](
	testName string, implementation func(InputStructure) OutputStructure, specialJudge func(InputStructure, OutputStructure, OutputStructure) bool,
	testCases ...SpecialJudgeTestCase[InputStructure, OutputStructure],
) TestCases[InputStructure, OutputStructure] {
	return &SpecialJudgeTestCases[InputStructure, OutputStructure]{
		name:          testName,
		implement:     implementation,
		judgeFunction: specialJudge,
		cases:         testCases,
	}
}
