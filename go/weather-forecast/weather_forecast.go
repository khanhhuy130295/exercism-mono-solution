// Package weather tools.
package weather

var (
	// CurrentCondition: this is condition of the current weather.
	CurrentCondition string
	//CurrentLocation: this is the current location.
	CurrentLocation string
)

// Forecast function return the string that's cotains the current location and the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
