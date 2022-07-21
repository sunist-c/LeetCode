// PruneTree: Leetcode 814 - Medium

package day3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PruneTree(root *TreeNode) *TreeNode {
	if root != nil {
		if root.Left != nil {
			root.Left = PruneTree(root.Left)
		}
		if root.Right != nil {
			root.Right = PruneTree(root.Right)
		}
		if root.Val == 0 && root.Left == nil && root.Right == nil {
			return nil
		}
	}
	return root
}
