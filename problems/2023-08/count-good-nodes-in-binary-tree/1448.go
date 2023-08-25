package count_good_nodes_in_binary_tree

func goodNodes(root *TreeNode) int {
	result := 0
	dfs(root, -10001, &result)
	return result
}

func dfs(root *TreeNode, currentMax int, result *int) {
	if root == nil {
		return
	}

	if root.Val >= currentMax {
		*result++
		currentMax = root.Val
	}

	dfs(root.Left, currentMax, result)
	dfs(root.Right, currentMax, result)
}
