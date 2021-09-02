package lasagna

func OvenTime() int {
	return 40
}

func RemainingOvenTime(minutes int) int {
	return OvenTime() - minutes
}

func PreparationTime(layers int) int {
	return 2 * layers
}

func ElapsedTime(layers, minutes int) int {
	return PreparationTime(layers) + minutes
}
