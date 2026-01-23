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
		// check the length of the string slice against the expected string slice
		// if they don't match use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("ERROR: Lengths not equal [%d] [%d]", len(actual), len(c.expected))
		}
		for i := range actual {
			// when iterating over the inner words in each case
			// check the result of running cleanInput against the input case against the expected test result
			// use i for this
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				// report error if the words do not match
				t.Errorf("Words do not match: [%v] [%v]", word, expectedWord)
			}
		}
	}
}
