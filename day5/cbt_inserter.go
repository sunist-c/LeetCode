// CBTInserter: Leetcode 919 - Medium

package day5

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BFS(t *TreeNode) *TreeNode {
	c := make(chan *TreeNode, 1000)
	defer close(c)

	c <- t
	for len(c) != 0 {
		node := <-c
		if node == nil {
			return nil
		}

		if node.Left == nil || node.Right == nil {
			return node
		} else {
			c <- node.Left
			c <- node.Right
		}
	}

	return nil
}

type CBTInserter struct {
	index *TreeNode
	root  *TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	tree := CBTInserter{
		index: BFS(root),
		root:  root,
	}

	return tree
}

func (c *CBTInserter) Insert(val int) int {
	c.index = BFS(c.root)
	fatherVal := c.index.Val
	node := &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}

	if c.index.Left == nil {
		c.index.Left = node
	} else {
		c.index.Right = node
	}

	return fatherVal
}

func (c *CBTInserter) GetRoot() *TreeNode {
	return c.root
}
