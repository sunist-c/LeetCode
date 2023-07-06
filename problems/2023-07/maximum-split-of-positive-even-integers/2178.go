package maximum_split_of_positive_even_integers

func maximumEvenSplit(finalSum int64) []int64 {
	if finalSum%2 != 0 {
		return []int64{}
	}

	var result []int64
	for currentEven := int64(2); currentEven <= finalSum; currentEven += 2 {
		result = append(result, currentEven)
		finalSum = finalSum - currentEven
	}

	result[len(result)-1] += finalSum

	return result
}
