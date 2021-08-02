package space

type Planet string

const Earth = 31_557_600.0

func years_during(factor float64, seconds float64) float64 {
	return seconds / (factor * Earth)
}

func Age(seconds float64, planet Planet) float64 {
	switch planet {
	case "Mercury":
		return years_during(0.240_846_7, seconds)
	case "Venus":
		return years_during(0.615_197_26, seconds)
	case "Earth":
		return years_during(1.0, seconds)
	case "Mars":
		return years_during(1.880_815_8, seconds)
	case "Jupiter":
		return years_during(11.862_615, seconds)
	case "Saturn":
		return years_during(29.447_498, seconds)
	case "Uranus":
		return years_during(84.016_846, seconds)
	case "Neptune":
		return years_during(164.79132, seconds)
	default:
		return float64(seconds) * 1.0
	}
}
