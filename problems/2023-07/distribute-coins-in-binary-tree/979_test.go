package distribute_coins_in_binary_tree

import (
	"leetcode/cases"
	"testing"
)

func TestDistributedCoinsInBinaryTree(t *testing.T) {
	testCases := []cases.CommonJudgeTestCase[*TreeNode, int]{
		{
			Input:      &TreeNode{Val: 3, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 0}},
			WantOutput: 2,
		},
		{
			Input:      &TreeNode{Val: 0, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 0}},
			WantOutput: 3,
		},
		{
			Input:      &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
			WantOutput: 2,
		},
		{
			Input:      &TreeNode{Val: 1, Left: &TreeNode{Val: 0, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 0}},
			WantOutput: 4,
		},
		{
			Input:      &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}},
			WantOutput: 0,
		},
	}

	t.Run(cases.NewCommonTestCases("DistributeCoinsInBinaryTree", distributeCoins, testCases...).JudgeFunction())
}
