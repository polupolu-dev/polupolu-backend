package validater

import (
	"errors"
)

// int の範囲
func MinMaxInt(n int, min int, max int) error {
	if n < min || n > max {
		return errors.New("out of range")
	}

	return nil
}
