package three_sum

import (
	"sort"
)

func threeSum(numbers []int) [][]int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	result, length := [][]int{}, len(numbers)

	for i, number := range numbers {
		if number > 0 {
			return result
		}

		left := i + 1
		right := length - 1

		for left < right {
			if number+numbers[left]+numbers[right] > 0 {
				right--
			} else if number+numbers[left]+numbers[right] < 0 {
				left++
			}
			if number+numbers[left]+numbers[right] == 0 {
				break
			}
		}

		if len(result) != 0 && left >= 0 && left < length && right >= 0 && right < length && left < right {
			results := result[len(result)-1]
			if number+numbers[left]+numbers[right] == 0 && number == results[0] && numbers[left] == results[1] && numbers[right] == results[2] && left < right {
				continue
			} else if number+numbers[left]+numbers[right] == 0 && left < right {
				result = append(result, []int{number, numbers[left], numbers[right]})
			}
		} else if left >= 0 && left < length && right >= 0 && right < length && number+numbers[left]+numbers[right] == 0 && left < right {
			result = append(result, []int{number, numbers[left], numbers[right]})
		}
	}

	return result
}
