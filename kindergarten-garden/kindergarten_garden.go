package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

type Cup rune

type CupPair [2]Cup

type ChildCups [2]CupPair

type Garden struct {
	flowers map[string][]string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	parts := strings.Split(diagram, "\n")
	//fmt.Printf("parts = %v, l = %d\n", parts, len(parts))
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid diagram")
	}
	cups := []Cup(parts[1])
	if len(cups)%2 != 0 {
		return nil, fmt.Errorf("invalid diagram, odd number of cups")
	}
	top := makePairs(cups)
	cups = []Cup(parts[2])
	if len(cups)%2 != 0 {
		return nil, fmt.Errorf("invalid diagram, odd number of cups")
	}
	bottom := makePairs(cups)
	if len(top) != len(bottom) {
		return nil, fmt.Errorf("invalid diagram, mismatch size")
	}

	garden := makeGarden(top, bottom)
	lookup := map[string][]string{}
	flowers := map[Cup]string{
		'G': "grass",
		'C': "clover",
		'R': "radishes",
		'V': "violets",
	}
	ch := []string{}
	ch = append(ch, children...)
	sort.Strings(ch)
	for i, child := range ch {
		if tr, ok := translateFlowers(flowers, garden[i]); !ok {
			return nil, fmt.Errorf("invalid flower code")
		} else if _, ok := lookup[child]; ok {
			return nil, fmt.Errorf("duplicated child")
		} else {
			lookup[child] = tr
		}
	}
	return &Garden{flowers: lookup}, nil
}

func makePairs(cups []Cup) []CupPair {
	result := []CupPair{}
	for i := 0; i < len(cups)-1; i += 2 {
		pair := CupPair{cups[i], cups[i+1]}
		result = append(result, pair)
	}
	return result
}

func makeGarden(top, bottom []CupPair) []ChildCups {
	result := []ChildCups{}
	for i := 0; i < len(top); i++ {
		cups := ChildCups{top[i], bottom[i]}
		result = append(result, cups)
	}
	return result
}

func translateFlowers(flowers map[Cup]string, garden ChildCups) ([]string, bool) {
	result := []string{}
	for _, pair := range garden {
		for _, cup := range pair {
			f, ok := flowers[cup]
			if !ok {
				return nil, false
			}
			result = append(result, f)
		}
	}
	return result, true
}

func (g *Garden) Plants(child string) ([]string, bool) {
	flowers, ok := g.flowers[child]
	return flowers, ok
}
