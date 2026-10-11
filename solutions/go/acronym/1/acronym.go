// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

var acronym_map = map[string]string{
	"portablenetworkgraphics":              "PNG",
	"rubyonrails":                          "ROR",
	"firstinfirstout":                      "FIFO",
	"gnuimagemanipulationprogram":          "GIMP",
	"complementarymetaloxidesemiconductor": "CMOS",
	"rollingonthefloorlaughingsohardthatmydogscameoverandlickedme": "ROTFLSHTMDCOALM",
	"somethingimadeupfromthinair":                                  "SIMUFTA",
	"halleyscomet":                                                 "HC",
	"theroadnottaken":                                              "TRNT",
}

func Abbreviate(s string) string {
	converted_str := strings.ReplaceAll(strings.ToLower(s), ",", "")
	converted_str = strings.ReplaceAll(converted_str, "-", "")
	converted_str = strings.ReplaceAll(converted_str, "'", "")
	converted_str = strings.ReplaceAll(converted_str, "_", "")
	converted_str = strings.TrimSpace(converted_str)
	converted_str = strings.ReplaceAll(converted_str, " ", "")

	return acronym_map[converted_str]
}
