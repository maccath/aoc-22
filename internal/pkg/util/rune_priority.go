package util

func GetPriority(c rune) int {
	// Lowercase
	if int(c) > 96 {
		return int(c) - 96
	}

	// Uppercase
	return int(c) - 38
}
