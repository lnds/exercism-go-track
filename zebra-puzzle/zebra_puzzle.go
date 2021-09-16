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
	for _, b := range permutations([]int{coffee, tea, orangejuice, water}) {
		beverages = append(beverages, []int{b[0], b[1], milk, b[2], b[3]})
	}

	// by 8: Kools are smoked in the yellow house.
	smokes := [][]int{}
	for _, s := range permutations([]int{oldgold, chesterfields, luckystrike, parliaments}) {
		ss := append([]int{kools}, s...)
		smokes = append(smokes, ss)
	}

	// by 12: Kools are smoked in the house next to the house where the horse is kept.
	pets := [][]int{}
	for _, p := range permutations([]int{dog, snails, fox, zebra}) {
		pets = append(pets, []int{p[0], horse, p[1], p[2], p[3]})
	}

	solution = Solution{DrinksWater: "", OwnsZebra: ""}
	iters := 0
	for _, c := range colors {
		for _, r := range residents {
			for _, b := range beverages {
				for _, s := range smokes {
					for _, p := range pets {
						houses := buildHouses(c, r, b, s, p)
						solution.DrinksWater = findResident(houses, beverage, water)
						solution.OwnsZebra = findResident(houses, pet, zebra)
						iters++
						if !c11(houses) || !any(houses, c2) || !any(houses, c3) || !any(houses, c4) ||
							!any(houses, c5) || !any(houses, c7) || !any(houses, c8) || !any(houses, c13) || !any(houses, c14) {
							continue
						}
						return
					}
				}
			}
		}
	}
	return
}

func buildHouses(colors, residents, beverages, smokes, pets []int) Houses {
	houses := Houses{}
	for i := first; i <= fifth; i++ {
		houses[i][color] = colors[i]
		houses[i][resident] = residents[i]
		houses[i][beverage] = beverages[i]
		houses[i][smoke] = smokes[i]
		houses[i][pet] = pets[i]
	}
	return houses
}

type Constraint [2]int // (atrribute, value)

func constraint(attribute1, value1, atrribute2, value2 int) [2]Constraint {
	return [2]Constraint{
		{attribute1, value1},
		{atrribute2, value2},
	}
}

// returns true if house satisfies all constraint
func satisfies(house House, constraints [2]Constraint) bool {
	atrribute1 := constraints[0][0]
	value1 := constraints[0][1]
	atrribute2 := constraints[1][0]
	value2 := constraints[1][1]
	return house[atrribute1] == value1 && house[atrribute2] == value2
}

// returns true if any house satisfies given constraints
func any(houses Houses, constraints [2]Constraint) bool {
	for _, house := range houses {
		if satisfies(house, constraints) {
			return true
		}
	}
	return false
}

var c2 = constraint(resident, englishman, color, red)           // 2
var c3 = constraint(resident, spaniard, pet, dog)               // 3
var c4 = constraint(beverage, coffee, color, green)             // 4
var c5 = constraint(resident, ukrainian, beverage, tea)         // 5
var c7 = constraint(smoke, oldgold, pet, snails)                // 7
var c8 = constraint(smoke, kools, color, yellow)                // 8
var c13 = constraint(smoke, luckystrike, beverage, orangejuice) // 13
var c14 = constraint(resident, japanese, smoke, parliaments)    // 14

func findResident(houses Houses, attribute, value int) string {
	for _, house := range houses {
		if house[attribute] == value {
			return nationality[house[resident]]
		}
	}
	return ""
}

// The man who smokes Chesterfields lives in the house next to the man with the fox.
func c11(houses Houses) bool {
	if houses[first][smoke] == chesterfields && houses[second][pet] == fox {
		return true
	}
	if houses[fifth][smoke] == chesterfields && houses[fourth][pet] == fox {
		return true
	}
	for i := second; i < fifth; i++ {
		if houses[i][smoke] == chesterfields {
			if houses[i-1][pet] == fox || houses[i+1][pet] == fox {
				return true
			}
		}
	}
	return false
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
