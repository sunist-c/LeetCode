package distribute_coins_in_binary_tree

func distributeCoins(root *TreeNode) int {
	var result = 0
	dfs(root, &result)
	return result
}

func dfs(root *TreeNode, counter *int) (steps, nodes int) {
	if root == nil {
		return 0, 0
	}

	leftSteps, leftNodes := dfs(root.Left, counter)
	rightSteps, rightNodes := dfs(root.Right, counter)

	steps, nodes = leftSteps+rightSteps+root.Val, leftNodes+rightNodes+1
	*counter += abs(steps - nodes)

	return steps, nodes
}

func abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
