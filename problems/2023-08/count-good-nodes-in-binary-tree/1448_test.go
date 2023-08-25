package count_good_nodes_in_binary_tree

import (
	"leetcode/cases"
	"testing"
)

func TestCountGoodNodesInBinaryTree(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[*TreeNode, int]{
		{
			Input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
					},
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
			},
			WantOutput: 4,
		},
		{
			Input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			WantOutput: 3,
		},
	}

	cases.NewCommonTestCases("CountGoodNodesInBinaryTree", goodNodes, testCases...).Run(t)
}
