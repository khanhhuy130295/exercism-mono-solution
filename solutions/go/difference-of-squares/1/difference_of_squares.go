package differenceofsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	if n <= 0 {
		return 0
	}
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

func Difference(n int) int {
	squareOfSum := SquareOfSum(n)
	sumOfSquares := SumOfSquares(n)

	return squareOfSum - sumOfSquares
}
