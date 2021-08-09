package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

type Robot struct {
	name string
}

func (robot *Robot) Name() (string, error) {
	if robot.name != "" {
		return robot.name, nil
	}

	if len(names) >= MAX_NAMES {
		return "", fmt.Errorf("no more names")
	}
	robot.name = randName()
	for names[robot.name] {
		robot.name = randName()
	}
	names[robot.name] = true
	return robot.name, nil
}

func (robot *Robot) Reset() {
	robot.name = ""
}

var letters = [26]rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

var names = map[string]bool{}
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

const MAX_PREFIX = 26 * 26
const MAX_SUFIX = 10 * 10 * 10
const MAX_NAMES = MAX_PREFIX * MAX_SUFIX

func randName() string {
	a := randUpperChar()
	b := randUpperChar()
	num := random.Intn(MAX_SUFIX)
	return fmt.Sprintf("%c%c%03d", a, b, num)
}

func randUpperChar() int {
	return 'A' + random.Intn(26)
}
