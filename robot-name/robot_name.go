package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

func (robot *Robot) Name() (string, error) {
	if robot.name == "" {
		robot.name = randName()
	}
	return robot.name, nil
}

func (robot *Robot) Reset() (string, error) {
	robot.name = randName()
	return robot.name, nil
}

var letters = [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

func randName() string {
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})
	return fmt.Sprintf("%c%c%03d", letters[0], letters[1], rand.Int()%1000)
}
