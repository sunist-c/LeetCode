package alternating_digit_sum

func alternateDigitSum(n int) int {
	number, sum, digital := n, 0, 0
	for n > 0 {
		digital += 1
		n /= 10
	}

	var signalMinus bool
	if digital%2 == 0 {
		signalMinus = true
	} else {
		signalMinus = false
	}

	for number > 0 {
		if signalMinus {
			sum -= number % 10
		} else {
			sum += number % 10
		}
		number /= 10
		signalMinus = !signalMinus
	}

	return sum
}
