package pangram

import "strings"

func IsPangram(input string) bool {
	var is_pangram bool = false
	seen := map[rune]bool{}

	for _, char := range strings.ToLower(input) {
		if char < 'a' || char > 'z' {
			continue
		}
		seen[char] = true
		if len(seen) == 26 {
			is_pangram = true
			break
		}
	}

	return is_pangram
}
