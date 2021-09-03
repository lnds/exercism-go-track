package partyrobot

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out his age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour string, direction string, distance float64) (result string) {
	result = fmt.Sprintf("Welcome to my party, %s!\n", name)
	result += fmt.Sprintf("You have been assigned to table %X.", table)
	result += fmt.Sprintf("Your table is %s, ", direction)
	result += fmt.Sprintf("exactly %.1f meters from here.\n", distance)
	result += fmt.Sprintf("You will be sitting next to %s", neighbour)
	return
}
