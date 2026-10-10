package isbnverifier

import (
	"strings"
)

func IsValidISBN(isbn string) bool {
	digits := strings.ReplaceAll(isbn, "-", "")
	if len(digits) != 10 {
		return false
	}

	sum := 0
	for i := 0; i < len(digits); i++ {
		digit := digits[i]
		value := 0

		if digit == 'X' && i == 9 {
			value = 10
		} else if digit >= '0' && digit <= '9' {
			value = int(digit - '0')
		} else {
			return false
		}

		sum += value * (10 - i)
	}

	return sum%11 == 0
}
