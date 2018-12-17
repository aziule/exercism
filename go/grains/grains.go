package grains

import (
	"errors"
)

const MaxUint64 = ^uint64(0)

var ErrInvalidInput = errors.New("Invalid number of inputs: must be between 1 and 64")

func Square(input int) (uint64, error) {
	if input < 1 || input > 64 {
		return 0, ErrInvalidInput
	}

	return 1 << uint(input-1), nil
}

func Total() uint64 {
	return MaxUint64
}
