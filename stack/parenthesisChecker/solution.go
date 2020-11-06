package parenthesisChecker

var (
	symbols = map[int32]int{
		'{': 1,
		'}': -1,
		'(': 2,
		')': -2,
		'[': 3,
		']': -3,
	}
)

// Can also be solved using bit vector
func IsBalanced(input string, allowInvertedInput bool) bool {
	if allowInvertedInput {
		return isBalancedIncludeInverted(input)
	} else {
		return isBalanced(input)
	}
}

func isBalanced(input string) bool {
	if input == "" {
		return true
	}

	dataStore := make([]int, 0)
	for _, letter := range input {
		current := symbols[letter]

		if len(dataStore) == 0 || current > 0 {
			dataStore = append(dataStore, current)
			continue
		}

		previous := dataStore[len(dataStore)-1]

		if previous > 0 && previous == (-1*current) {
			dataStore = dataStore[:len(dataStore)-1]
			continue
		}

		return false
	}

	return len(dataStore) == 0
}

func isBalancedIncludeInverted(input string) bool {
	if input == "" {
		return true
	}

	dataStore := make([]int, 0)
	for _, letter := range input {
		current := symbols[letter]

		if len(dataStore) == 0 {
			dataStore = append(dataStore, current)
			continue
		}

		previous := dataStore[len(dataStore)-1]

		if previous == (-1 * current) {
			dataStore = dataStore[:len(dataStore)-1]
			continue
		}

		dataStore = append(dataStore, current)
	}

	return len(dataStore) == 0
}
