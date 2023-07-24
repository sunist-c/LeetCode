package prime_number_of_set_bits_in_binary_representation

import "math/bits"

func countPrimeSetBits(left int, right int) (result int) {
	result = 0
	for lower, upper := left, right; lower <= upper; lower++ {
		count := 665772 >> bits.OnesCount(uint(left)) & 1
		result += count
	}
	return result
}
