package darts

func Score(x, y float64) int {

	var distance_square = x*x + y*y
	// center circle radius = 1 * 1
	if distance_square <= 1 {
		return 10
	}
	// middle circle radius = 5 * 5
	if distance_square <= 25 {
		return 5
	}
	// outer circle radius = 10 * 10
	if distance_square <= 100 {
		return 1
	}

	return 0
}
