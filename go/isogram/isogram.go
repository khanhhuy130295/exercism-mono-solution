package isogram

import "strings"

func IsIsogram(word string) bool {
	var isograms bool = true

	seen := map[rune]bool{}

	for _, char := range strings.ToLower(word) {
		if char == ' ' || char == '-' {
			continue
		}

		if seen[char] {
			return false

		}
		seen[char] = true

	}

	return isograms
}
