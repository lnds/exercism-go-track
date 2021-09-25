package bowling

import "fmt"

type State int

const (
	notPlayed = iota
	waiting
	open
	spare
	strike
)

type Frame struct {
	slots []int
	state State
	left  int
}

func (f Frame) roll(pins int) (Frame, error) {
	switch {
	case f.state == notPlayed && pins < 10:
		return Frame{slots: []int{pins}, state: waiting, left: 10 - pins}, nil
	case f.state == notPlayed && pins == 10:
		return Frame{slots: []int{pins}, state: strike, left: 0}, nil
	case f.state == waiting && pins+f.slots[0] == 10:
		return Frame{slots: []int{f.slots[0], pins}, state: spare, left: 0}, nil
	case f.state == waiting && pins+f.slots[0] < 10:
		return Frame{slots: []int{f.slots[0], pins}, state: open, left: 10 - f.slots[0] - pins}, nil
	default:
		return Frame{}, fmt.Errorf("pin count exceeds pins on the lane")
	}
}

type Game struct {
	frame  int
	frames []Frame
}

func NewGame() *Game {
	frames := make([]Frame, 12)
	for i := range frames {
		frames[i].state = notPlayed
		frames[i].slots = []int{}
		frames[i].left = 0
	}
	return &Game{
		frame:  0,
		frames: frames,
	}
}

func (g *Game) Score() (int, error) {
	if !g.ended() {
		return 0, fmt.Errorf("score cannot taken until the end of the game")
	}
	sum := 0
	for i := 0; i < 10; i++ {
		frames := g.frames[i : i+3]
		sum += calcScore(frames)
	}
	return sum, nil
}

func calcScore(frames []Frame) int {
	switch frames[0].state {
	case strike:
		score := 10
		if len(frames[1].slots) > 0 {
			score += frames[1].slots[0]
		}
		if len(frames[1].slots) == 2 {
			score += frames[1].slots[1]
		} else {
			if len(frames[2].slots) > 0 {
				score += frames[2].slots[0]
			}
		}
		return score
	case spare:
		score := 10
		if len(frames[1].slots) > 0 {
			score += frames[1].slots[0]
		}
		return score
	case open:
		return frames[0].slots[0] + frames[0].slots[1]
	default:
		return 0
	}
}

func (g *Game) Roll(pins int) error {
	if pins < 0 {
		return fmt.Errorf("negative roll is invalid")
	}
	if g.ended() {
		return fmt.Errorf("can't roll after game over")
	}
	frame := g.frames[g.frame]
	frame, err := frame.roll(pins)
	if err != nil {
		return err
	}
	g.update(frame)
	return nil
}

func (g *Game) update(frame Frame) {
	g.frames[g.frame] = frame
	switch frame.state {
	case open, strike, spare:
		g.frame++
	}
}

func (g *Game) ended() bool {
	a := g.frames[9].state
	b := g.frames[10].state
	c := g.frames[11].state
	switch a {
	case strike:
		return (b == strike && c != notPlayed) || (b == open || b == spare)
	case spare:
		return b != notPlayed
	case open:
		return true
	default:
		return false
	}
}
