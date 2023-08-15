package find_and_replace_in_string

import "strings"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	replace := map[int]string{}
	replaceLength := map[int]int{}

	for i := 0; i < len(indices); i++ {
		if strings.HasPrefix(s[indices[i]:], sources[i]) {
			replace[indices[i]] = targets[i]
			replaceLength[indices[i]] = len(sources[i])
		}
	}

	stringBuilder := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if v, ok := replace[i]; ok {
			stringBuilder.WriteString(v)
			i += replaceLength[i] - 1
		} else {
			stringBuilder.WriteByte(s[i])
		}
	}

	return stringBuilder.String()
}
