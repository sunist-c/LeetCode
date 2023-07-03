package add_two_numbers_ii

import (
	"strconv"
)

func reverseBytes(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

func calBig(num1 string, num2 string) string {
	nb1, nb2 := []byte(num1), []byte(num2)
	if len(nb1) < len(nb2) {
		nb1, nb2 = nb2, nb1
	}

	sum := byte(0)
	for i, j := len(nb1)-1, len(nb2)-1; i >= 0; i, sum = i-1, sum/10 {
		if j >= 0 {
			sum += nb2[j] - '0'
			j--
		}
		sum += nb1[i] - '0'
		nb1[i] = (sum % 10) + '0'
	}
	if sum != 0 {
		nb1 = append([]byte{'1'}, nb1...)
	}
	return reverseBytes(string(nb1))
}

func convertLinkedListToString(linkedList *ListNode) (result string) {
	currentPointer, resultString := linkedList, ""
	for currentPointer != nil {
		resultString += strconv.Itoa(currentPointer.Val)
		currentPointer = currentPointer.Next
	}

	if resultString == "" {
		return "0"
	} else {
		return resultString
	}
}

func convertStringToLinkedList(numberString string) *ListNode {
	var headPointer = &ListNode{
		Val:  0,
		Next: nil,
	}

	for _, i := range numberString {
		newNodeNumber, _ := strconv.Atoi(string(i))
		newNode := &ListNode{
			Val:  newNodeNumber,
			Next: headPointer.Next,
		}
		headPointer.Next = newNode
	}

	if headPointer.Next == nil {
		return &ListNode{
			Val:  0,
			Next: nil,
		}
	}

	return headPointer.Next
}

func addTwoNumbers(list1 *ListNode, list2 *ListNode) *ListNode {
	number1, number2 := convertLinkedListToString(list1), convertLinkedListToString(list2)
	return convertStringToLinkedList(calBig(number1, number2))
}
