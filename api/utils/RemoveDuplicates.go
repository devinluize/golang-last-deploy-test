package utils

func DistinctIntegers(input []int) []int {
	// A map to keep track of unique integers.
	unique := make(map[int]bool)
	var result []int

	for _, value := range input {
		if _, found := unique[value]; !found {
			unique[value] = true
			result = append(result, value)
		}
	}

	return result
}
