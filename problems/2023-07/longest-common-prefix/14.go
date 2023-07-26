package longest_common_prefix

func longestCommonPrefix(strs []string) string {
	var i = 0
	str := strs[findMaxLengthString(strs)]
	for ; i < len(str); i++ {
		for _, s := range strs {
			if i >= len(s) {
				return str[:i]
			}
			if s[i] != str[i] {
				return str[:i]
			}
		}
	}
	return str
}

func findMaxLengthString(strings []string) (index int) {
	maxLength := 0
	for i, str := range strings {
		if len(str) > maxLength {
			maxLength = len(str)
			index = i
		}
	}
	return index
}
