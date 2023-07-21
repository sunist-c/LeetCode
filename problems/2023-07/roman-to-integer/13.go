package roman_to_integer

import "strings"

func romanToInt(s string) int {
	s = strings.ReplaceAll(s, "IV", "F")
	s = strings.ReplaceAll(s, "IX", "S")
	s = strings.ReplaceAll(s, "XL", "T")
	s = strings.ReplaceAll(s, "XC", "U")
	s = strings.ReplaceAll(s, "CD", "O")
	s = strings.ReplaceAll(s, "CM", "P")
	runes := []roman(s)
	var result int
	for i := 0; i < len(runes); i++ {
		result += runes[i].value()
	}

	return result
}

type roman rune

func (r roman) value() int {
	switch r {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	case 'F':
		return 4
	case 'S':
		return 9
	case 'T':
		return 40
	case 'U':
		return 90
	case 'O':
		return 400
	case 'P':
		return 900
	default:
		return 0
	}
}
