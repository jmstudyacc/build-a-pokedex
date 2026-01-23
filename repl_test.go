package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		// first case
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// second case
		{
			input:    "I choose you Pikachu!",
			expected: []string{"i", "choose", "you", "pikachu!"},
		},
		// third case
		{
			input:    "123",
			expected: []string{"123"},
		},
		// fourth case
		{
			input:    "!!! £££ $$$ !£$",
			expected: []string{"!!!", "£££", "$$$", "!£$"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// check the length of the slice against the expected slice
		// if they don't match use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("ERROR: Lengths not equal [%d] [%d]", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Words do not match: [%v] [%v]", word, expectedWord)
			}
		}
	}
}
