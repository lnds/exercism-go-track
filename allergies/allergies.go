package allergies

var Allergens []string = []string{
	"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats",
}

func Allergies(score uint) []string {
	allergies := []string{}
	for i, a := range Allergens {
		if score&(1<<i) != 0 {
			allergies = append(allergies, a)
		}
	}
	return allergies
}

func AllergicTo(score uint, allergy string) bool {
	allergies := Allergies(score)
	for _, a := range allergies {
		if allergy == a {
			return true
		}
	}
	return false
}
