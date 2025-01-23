package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Charmander is cool",
			expected: []string{"charmander", "is", "cool"},
		},
		{
			input: "It's a Jigglypuff as seen from above",
			expected: []string{"it's", "a", "jigglypuff", "as", "seen", "from", "above"},
		},
	}

	passCount := 0
	failCount := 0

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Mismatch in output length. Input: %v \nExpected: %v\nActual: %v", c.input, len(c.expected), len(actual))

			// Count the fail & continue if they don't match so that we don't panic while looping
			failCount++
			continue
		}

		// If any of the words fail, it should only count as 1 failed test
		failed := false

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Mismatch in words. Input: %v\nExpected: %v\nActual: %v", c.input, expectedWord, word)
				failed = true
			}
		}

		if failed {
			failCount++
		} else {
			passCount++
		}
	}
	
	t.Logf("\nSuccess: %v\nFailure: %v", passCount, failCount)	
}