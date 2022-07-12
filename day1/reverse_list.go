package day1

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	// empty list
	if head == nil || head.Next == nil {
		return head
	}

	// init
	operate, node, prev := head, head.Next, head

	// action
	for operate != nil {
		if prev == operate {
			operate.Next = nil
		} else {
			operate.Next = prev
		}

		prev = operate
		operate = node
		if node != nil {
			node = node.Next
		}
	}

	return prev
}
