package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	var result []string

	for _, char := range candidates {
		var is_anagrams bool = checkAnagrams(strings.ToLower(subject), strings.ToLower(char))
		if is_anagrams {
			result = append(result, char)
		}
	}
	return result
}

func checkAnagrams(str1 string, str2 string) bool {
	if str1 == str2 {
		return false
	}

	seen := map[rune]int{}

	for _, char := range strings.ToLower(str1) {
		seen[char] += 1
	}
	for _, char := range strings.ToLower(str2) {
		seen[char] -= 1
		if seen[char] < 0 {
			return false
		}
	}
	for _, count := range seen {
		if count != 0 {
			return false
		}

	}
	return true
}
