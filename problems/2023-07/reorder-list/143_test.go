package reorder_list

import (
	"leetcode/cases"
	"testing"
)

func TestReorderList(t *testing.T) {
	inputList1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}}
	outputList1 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}}

	inputList2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	outputList2 := &ListNode{Val: 1, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}}}

	inputList3 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	outputList3 := &ListNode{Val: 1, Next: &ListNode{Val: 2}}

	inputList4 := &ListNode{Val: 1}
	outputList4 := &ListNode{Val: 1}

	judgeFunction := func(input, output, wantOutput *ListNode) bool {
		for output != nil && wantOutput != nil {
			if output.Val != wantOutput.Val {
				return false
			}
			output = output.Next
			wantOutput = wantOutput.Next
		}

		return true
	}

	originFunction := func(input *ListNode) (output *ListNode) {
		reorderList(input)
		return input
	}

	testCases := []cases.SpecialJudgeTestCase[*ListNode, *ListNode]{
		{
			Input:      inputList1,
			WantOutput: outputList1,
		},
		{
			Input:      inputList2,
			WantOutput: outputList2,
		},
		{
			Input:      inputList3,
			WantOutput: outputList3,
		},
		{
			Input:      inputList4,
			WantOutput: outputList4,
		},
	}

	cases.NewSpecialJudgeTestCases("TestReorderList", originFunction, judgeFunction, testCases...).Run(t)
}
