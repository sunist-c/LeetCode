package serialize_and_deserialize_bst

import (
	"leetcode/cases"
	"testing"
)

func TestSerializeAndDeserializeBst(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[*TreeNode, struct{}]{
		{
			Input: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
		},
		{
			Input: &TreeNode{
				Val: 1,
			},
		},
	}

	judgeFunction := func(input *TreeNode, output, want struct{}) bool {
		var stepJudgeFunc func(input *TreeNode, output *TreeNode) bool
		stepJudgeFunc = func(input *TreeNode, output *TreeNode) bool {
			if input == nil && output == nil {
				return true
			}

			if input.Val != output.Val {
				return false
			}

			if !stepJudgeFunc(input.Left, output.Left) {
				return false
			}

			if !stepJudgeFunc(input.Right, output.Right) {
				return false
			}

			return true
		}

		ser := Constructor()
		deser := Constructor()
		tree := ser.serialize(input)
		ans := deser.deserialize(tree)
		return stepJudgeFunc(input, ans)
	}

	originFunction := func(input *TreeNode) struct{} {
		return struct{}{}
	}

	cases.NewSpecialJudgeTestCases("SerializeAndDeserializeBst", originFunction, judgeFunction, testCases...).Run(t)
}
