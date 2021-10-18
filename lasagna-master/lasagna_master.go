package lasagna

func PreparationTime(layers []string, timePerLayer int) int {
	if timePerLayer == 0 {
		timePerLayer = 2
	}
	return len(layers) * timePerLayer
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return noodles, sauce
}

func AddSecretIngredient(friendsList, myList []string) []string {
	result := []string{}
	result = append(result, myList...)
	result = append(result, friendsList[len(friendsList)-1])
	return result
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	factor := float64(portions) / 2.0
	result := []float64{}
	for _, q := range quantities {
		result = append(result, q*factor)
	}
	return result
}
