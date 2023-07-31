package reorder_list

func reorderList(head *ListNode) {
	var list []*ListNode
	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	for i := 0; i < len(list)/2; i++ {
		list[i].Next = list[len(list)-1-i]
		list[len(list)-1-i].Next = list[i+1]
	}

	if len(list) > 0 {
		list[len(list)/2].Next = nil
	}
}
