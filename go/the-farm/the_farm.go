package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amt, err := weightFodder.FodderAmount()
	switch {
	case err != nil && err != ErrScaleMalfunction:
		return 0, err
	case amt < 0:
		return 0, errors.New("negative fodder")
	case cows == 0:
		return 0, errors.New("division by zero")
	case cows < 0:
		return 0, &SillyNephewError{cows}
	}
	if err != nil {
		amt *= 2
	}
	return amt / float64(cows), nil
}
