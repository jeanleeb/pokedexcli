package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "pokemon",
			expected: []string{"pokemon"},
		},
		{
			input:    "  multiple   spaces  between words  ",
			expected: []string{"multiple", "spaces", "between", "words"},
		},
		{
			input:    "   leading and trailing spaces   ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input:    "MIXED Case InPUT",
			expected: []string{"mixed", "case", "input"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Expected word %d to be '%s', but got '%s'", i, expectedWord, word)
			}
		}
	}
}

func Test_cleanInput(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		text string
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cleanInput(tt.text)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("cleanInput() = %v, want %v", got, tt.want)
			}
		})
	}
}
