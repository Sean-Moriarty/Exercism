// Package weather returns the weather forcast.
package weather

// CurrentCondition is the current weather conditions.
var CurrentCondition string
// CurrentLocation is the current location for forecast.
var CurrentLocation string

// Forecast returns forcast for the city and condition passed.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
