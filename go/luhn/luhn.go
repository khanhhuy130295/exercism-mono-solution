package luhn

import (
	"strings"
)

func Valid(id string) bool {
	digits := strings.ReplaceAll(id, " ", "")
	if len(digits) <= 1 {
		return false
	}

	sum := 0
	double := false

	for i := len(digits) - 1; i >= 0; i-- {
		digit := digits[i]
		if digit < '0' || digit > '9' {
			return false
		}

		value := int(digit - '0')
		if double {
			value *= 2
			if value > 9 {
				value -= 9
			}
		}

		sum += value
		double = !double
	}

	return sum%10 == 0
}
