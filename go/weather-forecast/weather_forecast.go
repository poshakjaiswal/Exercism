// Package weather Tell about the weather.
package weather

// CurrentCondition reperesents the current location weather temperature.
var CurrentCondition string

// CurrentLocation reperesents the current location weather temperature.
var CurrentLocation string

// Forecast returns a string value equal telling the current weather conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
