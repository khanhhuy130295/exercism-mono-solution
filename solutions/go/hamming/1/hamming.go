package hamming

import "fmt"

func Distance(a, b string) (int, error) {
	var distance int = 0
	if a == "" && b == "" {
		return distance, nil
	}

	if len(a) != len(b) {
		return distance, fmt.Errorf("strands must be of equal length")
	}

	array_a := []rune(a)
	array_b := []rune(b)

	for i := 0; i < len(array_a); i++ {
		var char_a string = string(array_a[i])
		var char_b string = string(array_b[i])

		if char_a != char_b {
			distance++
		}
	}

	return distance, nil
}
