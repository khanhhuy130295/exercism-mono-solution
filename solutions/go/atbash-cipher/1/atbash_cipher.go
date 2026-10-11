package atbashcipher

import "strings"

func atbashLetter(letter rune) rune {
	if letter >= 'a' && letter <= 'z' {
		return 'z' - (letter - 'a')
	}

	return letter
}

func Atbash(s string) string {
	var result strings.Builder
	var count int = 0

	for _, char := range strings.ToLower(s) {
		isLetter := char >= 'a' && char <= 'z'
		isNumber := char >= '0' && char <= '9'

		if !isLetter && !isNumber {
			continue
		}
		if count > 0 && count%5 == 0 {
			result.WriteRune(' ')
		}
		if isLetter {
			char = atbashLetter(char)
		}
		result.WriteRune(char)
		count++
	}
	return result.String()
}
