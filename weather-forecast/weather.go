// Package weather log the weather condition for a city
package weather

var (
	// stores current weather condition
	CurrentCondition string
	// stores current location
	CurrentLocation string
)

// Log stores city and condition on global variables
func Log(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
