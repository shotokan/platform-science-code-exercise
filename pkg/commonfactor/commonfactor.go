package commonfactor

// HaveCommonFactors check if two numbers have a common factor besides one
//
// Args:
//     a: integer to compare
//     b: integer to compare
//
// Returns:
//     returns true or false depending on whether the numbers have
//     a common factor other than 1.
//
func HaveCommonFactors(a, b int) bool {
	smaller := a
	if b < smaller {
		smaller = b
	}

	commonFactor := false
	for i := 2; i <= smaller; i++ {
		if a%i == 0 && b%i == 0 {
			commonFactor = true
			break
		}
	}

	return commonFactor
}
