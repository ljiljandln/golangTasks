package main

import (
	"strings"
	"testing"
)

const (
	path = "in.txt"
)

func TestTask(t *testing.T) {
	tests := []struct {
		name, pattern                 string
		after, before, context        int
		count, ignore, invert, number bool
		expected                      []string
	}{
		{"TestAfter", "баила",
			1, 0, 0,
			false, false, false, false,
			[]string{"Не те сказки нам мать баила в колыбели!", "Иные песни мы в армии запевали!"}},

		{"TestBefore", "баила",
			0, 1, 0,
			false, false, false, false,
			[]string{"Чуждо нам это!", "Не те сказки нам мать баила в колыбели!"}},

		{"TestContext", "ссыпем",
			0, 0, 3,
			false, false, false, false,
			[]string{"заработаем,", "в махорку ссыпем,", "да в перерыве закурим!"}},

		{"TestCount", "не",
			0, 0, 0,
			true, false, false, false,
			[]string{"8"}},

		{"TestCountWithIgnore", "не",
			0, 0, 0,
			true, true, false, false,
			[]string{"10"}},

		{"TestCountWithIgnoreAndInvert", "не",
			0, 0, 0,
			true, true, true, false,
			[]string{"24"}},

		{"TestCountWithNumber", "фонит",
			0, 0, 0,
			false, false, false, true,
			[]string{"8: Тут кедром фонит по низам!"}},
	}

	for _, test := range tests {
		fl := Flags{path, test.pattern, test.after, test.before, test.context,
			test.count, test.ignore, test.invert, test.number}
		actual, _ := solver(&fl)
		if !isEqual(test.expected, actual) {
			t.Errorf("Result are not are expected on test %s\nexpected: %s\nactual: %s ",
				test.name, test.expected, actual)
		}
	}
}

func isEqual(expected, actual []string) bool {
	if len(expected) != len(actual) {
		return false
	}
	for i := range actual {
		if !strings.EqualFold(actual[i], expected[i]) {
			return false
		}
	}
	return true
}
