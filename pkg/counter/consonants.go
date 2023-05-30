package counter

// Consonants Count the number of consonants in a string
//
// Args:
//     str: a string that may or may not contain consonants.
//
// Returns:
//     returns the number of consonants found in the string.
//
func Consonants(str string) int {
	consonantCount := 0

	for _, char := range str {
		if ((char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')) && (!IsVowel(char)) {
			// counts just consonants
			consonantCount++
		}
	}

	return consonantCount
}
