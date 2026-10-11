// Package acronym converts phrases into initial-letter abbreviations.
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate returns the uppercase initial of each word in s.
//
// Spaces, hyphens, and underscores separate words. Other punctuation is
// ignored, so an apostrophe in "Halley's" does not create another word.
func Abbreviate(s string) string {
	var acronym strings.Builder
	acronym.Grow(len(s))

	atWordStart := true
	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			if atWordStart {
				acronym.WriteRune(unicode.ToUpper(r))
			}
			atWordStart = false
		case unicode.IsSpace(r) || r == '-' || r == '_':
			atWordStart = true
		}
	}

	return acronym.String()
}
