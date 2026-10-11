package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	a := []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
	return a
}

var Color_map = map[int]int{
	0: 0,
	1: 1,
	2: 2,
	3: 3,
	4: 4,
	5: 5,
	6: 6,
	7: 7,
	8: 8,
	9: 9,
}

func indexOf(items []string, target string) int {
	for i, item := range items {
		if item == target {
			return i
		}
	}
	return -1
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	var list = Colors()
	var idx = indexOf(list, color)
	return Color_map[idx]
}
