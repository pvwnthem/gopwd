package audit

func CheckDuplicates(passwords map[string]string) (string, string, bool) {
	// Create a reverse map to keep track of values and their corresponding keys
	reverseMap := make(map[string]string)
	var firstDuplicate, secondDuplicate string
	foundDuplicates := false

	for key, value := range passwords {
		if existingKey, exists := reverseMap[value]; exists {
			// Found a duplicate value
			firstDuplicate = existingKey
			secondDuplicate = key
			foundDuplicates = true
			break
		}
		reverseMap[value] = key
	}

	return firstDuplicate, secondDuplicate, foundDuplicates
}
