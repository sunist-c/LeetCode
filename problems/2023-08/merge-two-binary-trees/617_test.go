package merge_two_binary_trees

import (
	"leetcode/cases"
	"testing"
)

type mergeTwoBinaryTreesInput struct {
	root1 *TreeNode
	root2 *TreeNode
}

func judgeFunction(input mergeTwoBinaryTreesInput, output, wantOutput *TreeNode) bool {
	if output == nil && wantOutput == nil {
		return true
	}
	if output == nil || wantOutput == nil {
		return false
	}
	if output.Val != wantOutput.Val {
		return false
	}
	return judgeFunction(input, output.Left, wantOutput.Left) && judgeFunction(input, output.Right, wantOutput.Right)
}

func TestMergeTwoBinaryTrees(t *testing.T) {
	tree1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	tree2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	want := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 5,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	testCases := []cases.SpecialJudgeTestCase[mergeTwoBinaryTreesInput, *TreeNode]{
		{
			Input:      mergeTwoBinaryTreesInput{root1: tree1, root2: tree2},
			WantOutput: want,
		},
	}

	originFunction := func(input mergeTwoBinaryTreesInput) *TreeNode {
		return mergeTrees(input.root1, input.root2)
	}

	cases.NewSpecialJudgeTestCases("MergeTwoBinaryTrees", originFunction, judgeFunction, testCases...).Run(t)
}
