package nextLargerElement

func FindNextLargerElements(input []uint) []int {
	results := make([]int, 0)
	if len(input) == 0 {
		return results
	}

	dataStore := []uint{input[0]}
	for index := 1; index < len(input); index++ {
		curr := input[index]
		top := dataStore[len(dataStore)-1]

		if curr <= top {
			dataStore = append(dataStore, curr)
			continue
		}

		for top <= curr {
			results = append(results, int(curr))
			dataStore = dataStore[:len(dataStore)-1]

			if len(dataStore) <= 0 {
				break
			}

			top = dataStore[len(dataStore)-1]
		}

		dataStore = append(dataStore, curr)
	}

	remainingDataStoreSize := len(dataStore)
	for index := 0; index < remainingDataStoreSize; index++ {
		results = append(results, -1)
	}

	return results
}
