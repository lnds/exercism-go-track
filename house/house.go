package house

import "fmt"

var subject = []string{
	"house that Jack built.",
	"malt",
	"rat",
	"cat",
	"dog",
	"cow with the crumpled horn",
	"maiden all forlorn",
	"man all tattered and torn",
	"priest all shaven and shorn",
	"rooster that crowed in the morn",
	"farmer sowing his corn",
	"horse and the hound and the horn",
}

var verb = []string{
	"lay in",
	"ate",
	"killed",
	"worried",
	"tossed",
	"milked",
	"kissed",
	"married",
	"woke",
	"kept",
	"belonged to",
}

func ThisIs(n int) string {
	return fmt.Sprintf("This is the %s", subject[n-1])
}

func That(n int) string {
	return fmt.Sprintf("that %s the %s", verb[n-1], subject[n-1])
}

func Verse(n int) string {
	verse := ThisIs(n)
	for i := n; i > 1; i-- {
		verse += "\n" + That(i-1)
	}
	return verse
}

func Song() string {
	song := Verse(1)
	for i := 2; i <= 12; i++ {
		song += "\n\n" + Verse(i)
	}
	return song
}
