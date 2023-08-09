package subtract_the_product_and_sum_of_digits_of_an_integer

func subtractProductAndSum(n int) int {
	number := n
	sum, product := 0, 1
	for number > 0 {
		bit := number % 10
		number = number / 10
		sum += bit
		product *= bit
	}

	return product - sum
}
