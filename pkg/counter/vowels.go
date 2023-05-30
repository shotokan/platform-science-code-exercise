package counter

import "strings"

// Vowels Count the number of vowels in a string
//
// Args:
//     str: a string that may or may not contain vowels.
//
// Returns:
//     returns the number of vowels found in the string.
//
func Vowels(str string) int {
	vowelCount := 0

	for _, char := range str {
		if IsVowel(char) {
			vowelCount++
		}
	}

	return vowelCount
}

func IsVowel(v rune) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, v)
}
