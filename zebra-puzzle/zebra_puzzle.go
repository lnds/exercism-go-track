package zebra

type HouseID int

const (
	first HouseID = iota
	second
	middle
	fourth
	fifth
)

const (
	resident = iota
	color
	beverage
	smoke
	pet
)

type House [5]int
type Houses [5]House

const (
	red = iota
	blue
	green
	ivory
	yellow
)

const (
	englishman = iota
	spaniard
	ukrainian
	norwegian
	japanese
)

var nationality = [5]string{"Englishman", "Spaniard", "Ukrainian", "Norwegian", "Japanese"}

const (
	coffee = iota
	tea
	milk
	orangejuice
	water
)

const (
	oldgold = iota
	kools
	chesterfields
	luckystrike
	parliaments
)

const (
	dog = iota
	snails
	fox
	horse
	zebra
)

type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

func SolvePuzzle() (solution Solution) {

	// by 15: The Norwegian lives next to the blue house.
	//    => {_,blue,_,_,_}
	// by 6 : The green house is immediately to the right of the ivory house.
	//    => {_, blue, ivory, green, _}, {_, blue, _, ivory, green}
	// by 10: The Norwegian lives in the first house.
	// by 2: The Englishman lives in the red house.
	//	  => {yellow, blue, ivory, green, red}, {yellow, blue, red, ivory, green}}
	colors := [][]int{
		{yellow, blue, ivory, green, red},
		{yellow, blue, red, ivory, green}}

	// by 10: The Norwegian lives in the first house.
	// by 2: The Englishman lives in the red house.
	residents := [][]int{}
	for _, r := range permutations([]int{spaniard, ukrainian, japanese}) {
		residents = append(residents, []int{norwegian, r[0], r[1], r[2], englishman})
		residents = append(residents, []int{norwegian, r[0], englishman, r[1], r[2]})
	}

	// by 9: Milk is drunk in the middle house.
	beverages := [][]int{}
	for _, b := range permutations([]int{tea, orangejuice, water}) {
		// by 4: Coffee is drunk in the green house.
		beverages = append(beverages, []int{b[0], b[1], milk, coffee, b[2]})
		beverages = append(beverages, []int{b[0], b[1], milk, b[2], coffee})
	}

	// by 8: Kools are smoked in the yellow house.
	smokes := [][]int{}
	for _, s := range permutations([]int{oldgold, chesterfields, luckystrike, parliaments}) {
		smokes = append(smokes, []int{kools, s[0], s[1], s[2], s[3]})
	}

	// by 12: Kools are smoked in the house next to the house where the horse is kept.
	pets := [][]int{}
	for _, p := range permutations([]int{dog, snails, fox, zebra}) {
		pets = append(pets, []int{p[0], horse, p[1], p[2], p[3]})
	}

	solution = Solution{DrinksWater: "", OwnsZebra: ""}
	for _, c := range colors {
		for _, r := range residents {
			for _, b := range beverages {
				for _, s := range smokes {
					for _, p := range pets {
						if houses, ok := buildHouses(c, r, b, s, p); ok {
							solution.DrinksWater = findResident(houses, beverage, water)
							solution.OwnsZebra = findResident(houses, pet, zebra)
							return
						}
					}
				}
			}
		}
	}
	return
}

func buildHouses(colors, residents, beverages, smokes, pets []int) (Houses, bool) {
	houses := Houses{}
	for i := first; i <= fifth; i++ {
		// by 3: The Spaniard owns the dog.
		if residents[i] == spaniard && pets[i] != dog {
			return houses, false
		}
		// by 5: The Ukrainian drinks tea.
		if residents[i] == ukrainian && beverages[i] != tea {
			return houses, false
		}
		// by 7: The Old Gold smoker owns snails.
		if smokes[i] == oldgold && pets[i] != snails {
			return houses, false
		}
		// by 13: The Lucky Strike smoker drinks orange juice.
		if smokes[i] == luckystrike && beverages[i] != orangejuice {
			return houses, false
		}
		// by 14: The Japanese smokes Parliaments.
		if residents[i] == japanese && smokes[i] != parliaments {
			return houses, false
		}
		// by 11: The man who smokes Chesterfields lives in the house next to the man with the fox.
		if smokes[i] == chesterfields {
			if i == 0 && pets[i+1] != fox {
				return houses, false
			}
			if i == 5 && pets[i-1] != fox {
				return houses, false
			}
			if pets[i-1] != fox && pets[i+1] == fox {
				return houses, false
			}
		}
		houses[i][color] = colors[i]
		houses[i][resident] = residents[i]
		houses[i][beverage] = beverages[i]
		houses[i][smoke] = smokes[i]
		houses[i][pet] = pets[i]
	}
	return houses, true
}

func findResident(houses Houses, attribute, value int) string {
	for _, house := range houses {
		if house[attribute] == value {
			return nationality[house[resident]]
		}
	}
	return ""
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
