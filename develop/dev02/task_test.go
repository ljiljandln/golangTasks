package main

import (
	"strings"
	"testing"
)

func TestTask(t *testing.T) {
	tests := []struct {
		name, str, expected string
		err                 error
	}{
		{name: "strWithoutDigits", str: "abcd", expected: "abcd", err: nil},
		{name: "strWithDigits", str: "a4bc2d5e", expected: "aaaabccddddde", err: nil},
		{name: "strEmpty", str: "", expected: "", err: nil},
		{name: "onlyDigitsInStr", str: "45", expected: "", err: IncorrectStringError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := unpackString(test.str)
			if err != test.err {
				t.Error("Неожиданная ошибка, Ожидали :", test.err, "Получили:", err)
			}

			if !strings.EqualFold(actual, test.expected) {
				t.Error("Функция выдала некорректный результат, ожидалось: ", test.expected, "Получили: ", actual,
					"Ошибка произошла в тесте %s", test.name)
			}
		})
	}
}
