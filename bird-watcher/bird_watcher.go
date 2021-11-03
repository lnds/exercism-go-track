package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) (total int) {
	for _, birds := range birdsPerDay {
		total += birds
	}
	return
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) (total int) {
	index := (week - 1) * 7
	for i := index; i < index+7; i++ {
		total += birdsPerDay[i]
	}
	return
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	birdsPerDay[0]++
	for i := 1; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i]++
		}
	}
	return birdsPerDay
}
