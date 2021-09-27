package poker

import (
	"fmt"
	"sort"
	"strings"
)

const A = 14

type Pair struct {
	rep  string
	hand Hand
}

func BestHand(rep []string) ([]string, error) {
	pairs := []Pair{}
	for _, r := range rep {
		hand, err := parseHand(r)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, Pair{rep: r, hand: hand})
	}
	if len(pairs) <= 1 {
		return rep, nil
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].hand.rank == pairs[j].hand.rank {
			return compareSameRank(pairs[i].hand, pairs[j].hand) > 0
		}
		return pairs[i].hand.rank > pairs[j].hand.rank
	})
	winner := []string{pairs[0].rep}
	for i := 1; i < len(pairs) && tie(pairs[i].hand, pairs[0].hand); i++ {
		winner = append(winner, pairs[i].rep)
	}
	return winner, nil
}

type Suit string
type Value int

type Card struct {
	value Value
	suit  Suit
}

const (
	Diamond = "♢"
	Heart   = "♡"
	Spade   = "♤"
	Club    = "♧"
)

const HandSize = 5

type Rank int

const (
	HighCard Rank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type Hand struct {
	deck   []Card
	rank   Rank
	groups map[Value]int
}

func parseHand(rep string) (hand Hand, err error) {
	parts := strings.Split(rep, " ")
	if len(parts) != 5 {
		err = fmt.Errorf("invalid hand format")
		return
	}
	var card Card
	cards := []Card{}
	for i := 0; i < HandSize; i++ {
		card, err = parseCard(parts[i])
		if err != nil {
			return
		}
		cards = append(cards, card)
	}
	sort.Slice(cards, func(i, j int) bool { return cards[i].value < cards[j].value })
	hand.deck = cards
	hand.groups = groupCards(cards)
	hand.rank = classify(hand.deck, hand.groups)
	return
}

func classify(cards []Card, valueMap map[Value]int) Rank {
	groups := []int{}
	for _, v := range valueMap {
		groups = append(groups, v)
	}
	sort.Ints(groups)
	switch {
	case len(groups) == 2 && groups[0] == 1 && groups[1] == 4:
		return FourOfAKind
	case len(groups) == 2 && groups[0] == 2 && groups[1] == 3:
		return FullHouse
	case len(groups) == 3 && groups[0] == 1 && groups[1] == 1 && groups[2] == 3:
		return ThreeOfAKind
	case len(groups) == 3 && groups[0] == 1 && groups[1] == 2 && groups[2] == 2:
		return TwoPair
	case len(groups) == 4 && groups[0] == 1 && groups[1] == 1 && groups[2] == 1 && groups[3] == 2:
		return OnePair
	case flush(cards) && straight(cards):
		return StraightFlush
	case straight(cards):
		return Straight
	case flush(cards):
		return Flush
	default:
		return HighCard
	}
}

func flush(cards []Card) bool {
	suit := cards[0].suit
	for i := 1; i < HandSize; i++ {
		if suit != cards[i].suit {
			return false
		}
	}
	return true
}

func straight(cards []Card) bool {
	result := true
	for i := 0; i < HandSize-1; i++ {
		result = result && cards[i].value == cards[i+1].value-1 || (cards[i].value == 5 && cards[i+1].value == A)
	}
	return result
}

func groupCards(cards []Card) map[Value]int {
	groups := map[Value]int{}
	for _, card := range cards {
		groups[card.value]++
	}

	return groups
}

func parseCard(rep string) (card Card, err error) {
	values := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	value := 0
	for i, v := range values {
		if strings.HasPrefix(rep, v) {
			value = 2 + i
			rep = rep[len(v):]
			break
		}
	}
	if value == 0 {
		err = fmt.Errorf("invalid value of card")
		return
	}
	suit := Suit(rep)
	if suit != Diamond && suit != Club && suit != Heart && suit != Spade {
		err = fmt.Errorf("invalid suit of card")
		return
	}
	card = Card{value: Value(value), suit: suit}
	return
}

func tie(hand1, hand2 Hand) bool {
	if hand1.rank != hand2.rank {
		return false
	}
	if valueOf(hand1) != valueOf(hand2) {
		return false
	}
	return compareSameRank(hand1, hand2) == 0
}

func compareSameRank(hand1, hand2 Hand) int {
	switch hand1.rank {
	case HighCard:
		diff := 0
		for i := HandSize - 1; i >= 0; i-- {
			diff = int(hand1.deck[i].value) - int(hand2.deck[i].value)
			if diff != 0 {
				break
			}
		}
		return diff
	case OnePair:
		p1 := Value(0)
		for k, v := range hand1.groups {
			if v == 2 {
				p1 = k
				break
			}
		}
		p2 := Value(0)
		for k, v := range hand2.groups {
			if v == 2 {
				p2 = k
				break
			}
		}
		return int(p1) - int(p2)
	case TwoPair:
		p1 := []int{}
		k1 := 0
		for k, v := range hand1.groups {
			if v == 2 {
				p1 = append(p1, int(k))
			} else {
				k1 = int(k)
			}
		}
		p2 := []int{}
		k2 := 0
		for k, v := range hand2.groups {
			if v == 2 {
				p2 = append(p2, int(k))
			} else {
				k2 = int(k)
			}
		}
		sort.Ints(p1)
		sort.Ints(p2)
		if p1[1] != p2[1] {
			return p1[1] - p2[1]
		}
		if p1[0] != p2[0] {
			return p1[0] - p2[0]
		}
		return k1 - k2
	case ThreeOfAKind:
		p1 := Value(0)
		k1 := Value(0)
		for k, v := range hand1.groups {
			if v == 3 {
				p1 = k
			} else if k > k1 {
				k1 = k
			}
		}
		p2 := Value(0)
		k2 := Value(0)
		for k, v := range hand2.groups {
			if v == 3 {
				p2 = k
			} else if k > k2 {
				k2 = k
			}
		}
		if p1 == p2 {
			return int(k1) - int(k2)
		}
		return int(p1) - int(p2)
	case FullHouse:
		t1 := Value(0)
		p1 := Value(0)
		k1 := Value(0)
		for k, v := range hand1.groups {
			if v == 3 {
				t1 = k
			} else if v == 2 {
				p1 = k
			} else {
				k1 = k
			}
		}
		t2 := Value(0)
		p2 := Value(0)
		k2 := Value(0)
		for k, v := range hand2.groups {
			if v == 3 {
				t2 = k
			} else if v == 2 {
				p2 = k
			} else {
				k2 = k
			}
		}
		if t1 != t2 {
			return int(t1) - int(t2)
		}
		if p1 != p2 {
			return int(p1) - int(p2)
		}
		return int(k1) - int(k2)
	case FourOfAKind:
		p1 := Value(0)
		k1 := Value(0)
		for k, v := range hand1.groups {
			if v == 4 {
				p1 = k
			} else {
				k1 = k
			}
		}
		p2 := Value(0)
		k2 := Value(0)
		for k, v := range hand2.groups {
			if v == 4 {
				p2 = k
			} else {
				k2 = k
			}
		}
		if p1 == p2 {
			return int(k1) - int(k2)
		}
		return int(p1) - int(p2)
	case Straight, StraightFlush:
		p1 := int(hand1.deck[HandSize-1].value)
		p2 := int(hand2.deck[HandSize-1].value)
		if p1 == A {
			p1 = int(hand1.deck[HandSize-2].value)
		}
		if p2 == A {
			p2 = int(hand2.deck[HandSize-2].value)
		}
		return p1 - p2
	}
	return 0
}

func valueOf(hand Hand) Value {
	switch hand.rank {
	case HighCard, Flush, StraightFlush:
		return hand.deck[HandSize-1].value
	case OnePair:
		val := Value(0)
		for k, v := range hand.groups {
			if v == 2 {
				val = k
			}
		}
		return val
	case TwoPair:
		val := Value(0)
		for k, v := range hand.groups {
			if v == 2 {
				if k > val {
					val = k
				}
			}
		}
		return val
	case ThreeOfAKind, FullHouse:
		val := Value(0)
		for k, v := range hand.groups {
			if v == 3 {
				if k > val {
					val = k
				}
			}
		}
		return val
	case FourOfAKind:
		val := Value(0)
		for k, v := range hand.groups {
			if v == 4 {
				if k > val {
					val = k
				}
			}
		}
		return val
	case Straight:
		val := hand.deck[HandSize-1].value
		if A == val {
			val = 1
		}
		return val
	}
	return Value(0)
}
