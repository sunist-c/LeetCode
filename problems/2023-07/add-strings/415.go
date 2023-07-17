package add_strings

func addStrings(num1, num2 string) string {
	num1Arr, num2Arr := []byte(num1), []byte(num2)
	num1Len, num2Len := len(num1Arr), len(num2Arr)
	if num1Len+num2Len == 0 {
		return "0"
	}

	minLength, maxLength := 0, 0
	num1IsLonger := false
	if num1Len < num2Len {
		minLength = num1Len
		maxLength = num2Len
		num1IsLonger = false
	} else {
		minLength = num2Len
		maxLength = num1Len
		num1IsLonger = true
	}

	result := make([]byte, maxLength+1)
	carry := 0
	for i := 0; i < maxLength; i++ {
		sum := carry
		if i < minLength {
			sum += int(num1Arr[num1Len-1-i]-'0') + int(num2Arr[num2Len-1-i]-'0')
		} else if num1IsLonger {
			sum += int(num1Arr[num1Len-1-i] - '0')
		} else {
			sum += int(num2Arr[num2Len-1-i] - '0')
		}

		carry = sum / 10
		result[maxLength-i] = byte(sum%10) + '0'
	}

	if carry > 0 {
		result[0] = '1'
		return string(result)
	}

	return string(result[1:])
}
