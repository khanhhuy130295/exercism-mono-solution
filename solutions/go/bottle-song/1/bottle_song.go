package bottlesong

import (
	"fmt"
	"strings"
)

var numberWords = []string{
	"no", "One", "Two", "Three", "Four", "Five",
	"Six", "Seven", "Eight", "Nine", "Ten",
}

func Recite(startBottles, takeDown int) []string {
	var lyrics []string

	for bottles := startBottles; bottles > startBottles-takeDown; bottles-- {
		if len(lyrics) > 0 {
			lyrics = append(lyrics, "")
		}

		word := numberWords[bottles]
		bottle := "bottles"
		if bottles == 1 {
			bottle = "bottle"
		}
		remainingBottle := "bottles"
		if bottles-1 == 1 {
			remainingBottle = "bottle"
		}

		verse := fmt.Sprintf("%s green %s hanging on the wall,", word, bottle)
		lyrics = append(lyrics,
			verse,
			verse,
			"And if one green bottle should accidentally fall,",
			fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(numberWords[bottles-1]), remainingBottle),
		)
	}

	return lyrics
}
