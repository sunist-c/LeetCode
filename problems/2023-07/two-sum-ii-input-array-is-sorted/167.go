package two_sum_ii_input_array_is_sorted

func twoSumAnother(numbers []int, target int) []int {
	indexStart, indexEnd := 0, len(numbers)-1
	length := len(numbers)
	for indexStart < length {
		if numbers[indexStart]+numbers[indexEnd] >= target {
			for indexEnd > indexStart && numbers[indexStart]+numbers[indexEnd] > target {
				indexEnd--
			}
			if numbers[indexStart]+numbers[indexEnd] == target {
				return []int{indexStart + 1, indexEnd + 1}
			}
		} else {
			for indexEnd < length-1 && numbers[indexStart]+numbers[indexEnd] < target {
				indexEnd++
			}
			if numbers[indexStart]+numbers[indexEnd] == target {
				return []int{indexStart + 1, indexEnd + 1}
			}
		}

		indexStart++
	}

	return []int{}
}

func twoSum(numbers []int, target int) []int {
	for index, number := range numbers {
		if anotherIndex, exist := findNumberIndex(numbers, target-number, index); exist {
			if index < anotherIndex {
				return []int{index + 1, anotherIndex + 1}
			} else {
				return []int{anotherIndex + 1, index + 1}
			}
		}
	}
	return []int{}
}

func findNumberIndex(numbers []int, target int, ignoreIndex int) (result int, exist bool) {
	numberResultArray := append(numbers[:ignoreIndex], numbers[ignoreIndex:]...)
	findResult, existNumber := findIndex(numberResultArray, target, 0, len(numbers)-1)

	if !existNumber {
		return findResult, existNumber
	} else if findResult == ignoreIndex {
		return findResult + 1, existNumber
	} else {
		return findResult, existNumber
	}
}

func findIndex(numbers []int, target int, lower, upper int) (result int, exist bool) {
	if upper-lower < 0 {
		return 0, false
	}

	mid := lower + (upper-lower)/2
	if numbers[mid] == target {
		return mid, true
	} else if numbers[mid] < target {
		return findIndex(numbers, target, mid+1, upper)
	} else {
		return findIndex(numbers, target, lower, mid-1)
	}
}
