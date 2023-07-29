package linked_list_cycle

import (
	"leetcode/cases"
	"testing"
)

func TestLinkedListCycle(t *testing.T) {
	inputList1 := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4}}}}
	inputList1.Next.Next.Next.Next = inputList1.Next
	inputList2 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	inputList2.Next.Next = inputList2
	inputList3 := &ListNode{Val: 1}

	testCases := []cases.CommonJudgeTestCase[*ListNode, bool]{
		{
			Input:      inputList1,
			WantOutput: true,
		},
		{
			Input:      inputList2,
			WantOutput: true,
		},
		{
			Input:      inputList3,
			WantOutput: false,
		},
		{
			Input:      nil,
			WantOutput: false,
		},
	}

	cases.NewCommonTestCases("TestLinkedListCycle", hasCycle, testCases...).Run(t)
}
