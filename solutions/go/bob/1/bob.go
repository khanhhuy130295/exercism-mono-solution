// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isUppercaseQuestion(sentence string) bool {
	hasLetter := false

	for _, char := range sentence {
		if unicode.IsLetter(char) {
			hasLetter = true
			if !unicode.IsUpper(char) {
				return false
			}
		}
	}

	return hasLetter && isQuestionMark(sentence)
}

func isQuestionMark(sentence string) bool {
	return strings.HasSuffix(strings.TrimSpace(sentence), "?")
}

func isUpperCaseWithoutQuestion(sentence string) bool {
	hasLetter := false

	for _, char := range sentence {
		if unicode.IsLetter(char) {
			hasLetter = true
			if !unicode.IsUpper(char) {
				return false
			}
		}
	}

	return hasLetter && !isQuestionMark(sentence)
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	if strings.TrimSpace(remark) == "" {
		return "Fine. Be that way!"
	}

	if isUppercaseQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	}

	if isUpperCaseWithoutQuestion(remark) {
		return "Whoa, chill out!"
	}

	if isQuestionMark(remark) {
		return "Sure."
	}

	return "Whatever."
}
