package element_appearing_more_than_25_in_sorted_array

func findSpecialInteger(arr []int) int {
	length := len(arr) / 4
	for i := 0; i < len(arr); i++ {
		if arr[i] == arr[i+length] {
			return arr[i]
		}
	}
	return 0
}
