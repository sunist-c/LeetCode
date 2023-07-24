package jewels_and_stones

func numJewelsInStones(jewels string, stones string) (number int) {
	jewelsMap := map[rune]struct{}{}
	for _, jewel := range jewels {
		jewelsMap[jewel] = struct{}{}
	}

	for _, stone := range stones {
		if _, exist := jewelsMap[stone]; exist {
			number++
		}
	}

	return number
}
