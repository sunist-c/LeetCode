package cases

import (
	"fmt"
	"testing"
)

type CommonJudgeTestCase[InputStructure any, OutputStructure any] struct {
	Input      InputStructure
	WantOutput OutputStructure
}

type CommonJudgeTestCases[InputStructure any, OutputStructure comparable] struct {
	name      string
	cases     []CommonJudgeTestCase[InputStructure, OutputStructure]
	implement func(InputStructure) OutputStructure
}

func (cases *CommonJudgeTestCases[InputStructure, OutputStructure]) JudgeFunction() (name string, testingFunction func(t *testing.T)) {
	return fmt.Sprintf("Test%s-CommonJudge", cases.name),
		func(t *testing.T) {
			stat := newStatistics()
			stat.start()

			for i, testCase := range cases.cases {
				t.Run(fmt.Sprintf("%s-%d", cases.name, i+1), func(t *testing.T) {
					output := cases.implement(testCase.Input)
					if output != testCase.WantOutput {
						t.Errorf("Failed to Run Test in %s: %d\n"+
							"Input: %#v\n"+
							"Want Output: %#v\n"+
							"Get Output: %#v\n",
							cases.name, i+1, testCase.Input, testCase.WantOutput, output,
						)
					}
				})
			}

			memoryAllocBytes, elapsed := stat.end()
			t.Logf("Task %s Finished Took %s(Total %d Cases) - Alloc %.4f MB Memory", cases.name, elapsed.String(), len(cases.cases), float64(memoryAllocBytes)/float64(1024*1024))
		}
}

func NewCommonTestCases[InputStructure any, OutputStructure comparable](
	testName string, implementation func(InputStructure) OutputStructure,
	testCases ...CommonJudgeTestCase[InputStructure, OutputStructure],
) TestCases[InputStructure, OutputStructure] {
	return &CommonJudgeTestCases[InputStructure, OutputStructure]{
		name:      testName,
		implement: implementation,
		cases:     testCases,
	}
}
