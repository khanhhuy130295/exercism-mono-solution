package cars

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var rate float64 = float64(successRate) / 100
	return float64(productionRate) * rate
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var rate float64 = float64(successRate) / 100
	return int(float64(productionRate)*rate) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	const base_cost_per_car int = 10000
	const group_of_ten_cost_per_car int = 95000

	var count_group_of_ten = carsCount / 10

	var count_individual = uint(carsCount - count_group_of_ten*10)

	var total uint = uint(count_group_of_ten)*uint(group_of_ten_cost_per_car) + uint(count_individual)*uint(base_cost_per_car)

	return total
}
