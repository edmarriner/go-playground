package go_playground

import "testing"

func TestValidWordsPass(t *testing.T) {
	cases := []string{
		"racecar",
		"noon",
	}

	for _, input := range cases {
		valid, err := IsStringPalindrome([]byte(input))

		if err != nil {
			t.Errorf("Failed to pass a valid word '%s', error: %e", input, err)
		}

		if !valid {
			t.Errorf("Failed to pass a valid word '%s'", input)
		}
	}
}

func TestIsCaseInsensitive(t *testing.T) {
	cases := []string{
		"Racecar",
		"NoOn",
	}

	for _, input := range cases {
		valid, err := IsStringPalindrome([]byte(input))

		if err != nil {
			t.Errorf("Failed to pass a valid word '%s', error: %e", input, err)
		}

		if !valid {
			t.Errorf("Failed to pass a valid word '%s'", input)
		}
	}
}

func TestPhrases(t *testing.T) {
	cases := []string{
		"a nut for a jar of tuna",
		"stressed desserts",
	}

	for _, input := range cases {
		valid, err := IsStringPalindrome([]byte(input))

		if err != nil {
			t.Errorf("Failed to pass a valid phrase '%s', error: %e", input, err)
		}

		if !valid {
			t.Errorf("Failed to pass a valid phrase '%s'", input)
		}
	}
}
