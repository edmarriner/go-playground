package palindrome

import (
	"bytes"
	"errors"
	"unicode"
)

func IsStringPalindrome(input []byte) (bool, error) {

	if len(input) == 0 {
		return false, errors.New("empty input supplied")
	}

	var reverse []byte
	var source []byte

	for i := 0; i < len(input); i++ {

		if input[i] == 32 {
			continue
		}

		sourceByte := byte(unicode.ToLower(rune(input[i])))
		source = append(source, sourceByte)
	}

	for i := len(input) - 1; i >= 0; i-- {

		if input[i] == 32 {
			continue
		}

		// lowercase the byte
		reverseByte := byte(unicode.ToLower(rune(input[i])))

		reverse = append(reverse, reverseByte)
	}

	if bytes.Compare(source, reverse) != 0 {
		return false, nil
	}

	return true, nil
}
