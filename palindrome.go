package go_playground

import (
	"bytes"
	"errors"
)

func IsStringPalindrome(input []byte) (bool, error) {

	if len(input) == 0 {
		return false, errors.New("empty input supplied")
	}

	var reverse []byte

	for i := len(input) - 1; i >= 0; i-- {
		reverse = append(reverse, input[i])
	}

	if bytes.Compare(input, reverse) != 0 {
		return false, nil
	}

	return true, nil
}
