package twelve

import (
	"fmt"
	"strings"
)

const days = 12

var ords = [days]string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = [days]string{
	"a Partridge",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

func Song() string {
	song := make([]string, days)
	for i := 0; i < len(song); i++ {
		song[i] = Verse(i + 1)
	}
	return strings.Join(song, "\n")
}

func giftVerse(n int) (verse string) {
	for i := n; i > 0; i-- {
		verse += gifts[i] + ", "
	}
	if n > 0 {
		verse += "and "
	}
	verse += gifts[0]
	return
}

func Verse(n int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s in a Pear Tree.", ords[n-1], giftVerse(n-1))
}
