package rnatranscription

import "strings"

func ToRNA(dna string) string {
	if len(dna) == 0 {
		return ""
	}
	var rna strings.Builder

	for _, char := range dna {
		switch char {
		case 'G':
			rna.WriteString("C")
		case 'C':
			rna.WriteString("G")
		case 'T':
			rna.WriteString("A")
		case 'A':
			rna.WriteString("U")
		}
	}
	return rna.String()
}
