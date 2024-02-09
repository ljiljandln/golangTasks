package main

import "testing"

func TestTask(t *testing.T) {
	tests := []struct {
		name     string
		fl       flags
		lines    []string
		expected []string
	}{
		{"TestWithAnotherDelimiter",
			flags{1, "ttt", false},
			[]string{"5ttt6", "7ttt8"},
			[]string{"5", "7"}},

		{"TestWithOnlyDelimiter",
			flags{1, " ", true},
			[]string{"5", "7 8"},
			[]string{"7"}},

		{"TestWithColumn",
			flags{2, "ttt", true},
			[]string{"5ttt6", "7ttt8"},
			[]string{"6", "8"}},
	}

	for _, test := range tests {
		actual := makeCut(&test.fl, test.lines)
		if !isLinesEqual(test.expected, actual) {
			t.Errorf("the lines are different in test %s, expected %v, actual: %v",
				test.name, test.expected, actual)
		}
	}
}

func isLinesEqual(actual, expected []string) bool {
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			return false
		}
	}
	return true
}
