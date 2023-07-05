package k_items_with_the_maximum_sum

func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	result := 0
	if k <= numOnes {
		return k
	} else {
		k -= numOnes
		result += numOnes
	}

	if k <= numZeros {
		return result
	} else {
		k -= numZeros
		return result - k
	}
}
