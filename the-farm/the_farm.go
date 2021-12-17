package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

type SillyNephewError struct {
	cows int
}

func NewSillyNephewError(cows int) *SillyNephewError {
	return &SillyNephewError{cows: cows}
}

func (s SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	amount, err := weightFodder.FodderAmount()

	if err != nil && err == ErrScaleMalfunction {
		amount *= 2
		err = nil
	}

	switch {
	case err != nil:
		return 0.0, err
	case amount < 0.0:
		return 0, errors.New("Negative fodder")

	case cows == 0:
		return 0.0, errors.New("Division by zero")

	case cows < 0:
		return 0, NewSillyNephewError(cows)
	default:
		return amount / float64(cows), nil
	}
}
