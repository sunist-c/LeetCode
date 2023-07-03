package add_two_numbers_ii

import (
	"leetcode/cases"
	"testing"
)

type addTwoNumbersIIInput struct {
	list1 *ListNode
	list2 *ListNode
}

type addTwoNumbersIIOutput struct {
	result *ListNode
}

func convertNumberArrayToList(numbers ...int) *ListNode {
	headPointer := &ListNode{
		Val:  0,
		Next: nil,
	}
	currentNode := headPointer

	for _, number := range numbers {
		newNode := &ListNode{
			Val:  number,
			Next: nil,
		}
		currentNode.Next = newNode
		currentNode = newNode
	}

	return headPointer.Next
}

func convertListToNumberArray(list *ListNode) (numbers []int) {
	currentNode := list

	resultArray := []int{}
	for currentNode != nil {
		resultArray = append(resultArray, currentNode.Val)
		currentNode = currentNode.Next
	}

	if len(resultArray) == 0 {
		return []int{0}
	} else {
		return resultArray
	}
}

func TestAddTwoNumbersII(t *testing.T) {
	testCases := []cases.SpecialJudgeTestCase[addTwoNumbersIIInput, addTwoNumbersIIOutput]{
		{
			Input: addTwoNumbersIIInput{
				list1: convertNumberArrayToList(2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9),
				list2: convertNumberArrayToList(5, 6, 4, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 2, 4, 3, 9, 9, 9, 9),
			},
			WantOutput: addTwoNumbersIIOutput{
				result: convertNumberArrayToList(8, 0, 7, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 6, 4, 8, 7, 2, 4, 3, 8),
			},
		},
		{
			Input: addTwoNumbersIIInput{
				list1: convertNumberArrayToList(7, 2, 4, 3),
				list2: convertNumberArrayToList(5, 6, 4),
			},
			WantOutput: addTwoNumbersIIOutput{
				result: convertNumberArrayToList(7, 8, 0, 7),
			},
		},
		{
			Input: addTwoNumbersIIInput{
				list1: convertNumberArrayToList(0),
				list2: convertNumberArrayToList(0),
			},
			WantOutput: addTwoNumbersIIOutput{
				result: convertNumberArrayToList(0),
			},
		},
		{
			Input: addTwoNumbersIIInput{
				list1: convertNumberArrayToList(2, 4, 3),
				list2: convertNumberArrayToList(5, 6, 4),
			},
			WantOutput: addTwoNumbersIIOutput{
				result: convertNumberArrayToList(8, 0, 7),
			},
		},
	}

	originFunction := func(input addTwoNumbersIIInput) (output addTwoNumbersIIOutput) {
		return addTwoNumbersIIOutput{
			result: addTwoNumbers(input.list1, input.list2),
		}
	}
	judgeFunction := func(input addTwoNumbersIIInput, output addTwoNumbersIIOutput, wantOutput addTwoNumbersIIOutput) bool {
		got, want := convertListToNumberArray(output.result), convertListToNumberArray(wantOutput.result)
		if len(got) != len(want) {
			return false
		} else {
			for i := 0; i < len(got); i++ {
				if got[i] != want[i] {
					return false
				}
			}
			return true
		}
	}

	t.Run(cases.NewSpecialJudgeTestCases("AddTwoNumbersII", originFunction, judgeFunction, testCases...).JudgeFunction())
}
