package cases

import "testing"

type TestCases[InputStructure any, OutputStructure any] interface {
	JudgeFunction() (name string, testingFunction func(t *testing.T))
}
