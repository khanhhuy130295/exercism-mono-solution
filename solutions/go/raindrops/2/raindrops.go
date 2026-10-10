package raindrops

import (
	"strconv"
	"strings"
)

func Convert(number int) string {
	factors := []struct {
		value int
		sound string
	}{
		{3, "Pling"},
		{5, "Plang"},
		{7, "Plong"},
	}

	var result strings.Builder
	for _, factor := range factors {
		if number%factor.value == 0 {
			result.WriteString(factor.sound)
		}
	}

	if result.Len() == 0 {
		return strconv.Itoa(number)
	}

	return result.String()
}
