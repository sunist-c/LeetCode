package move_pieces_to_obtain_a_string

func canChange(start string, target string) bool {
	length, i, j := len(start), 0, 0
	for {
		for i < length && start[i] == '_' {
			i++
		}
		for j < length && target[j] == '_' {
			j++
		}
		if i == length && j == length {
			return true
		}
		if i == length || j == length || start[i] != target[j] {
			return false
		}
		if start[i] == 'L' && i < j {
			return false
		}
		if start[i] == 'R' && i > j {
			return false
		}
		i, j = i+1, j+1
	}
}
