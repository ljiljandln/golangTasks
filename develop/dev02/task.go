package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

var IncorrectStringError = errors.New("string without letters can't be unpacked")

func unpackString(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}
	runes := []rune(str)
	res := make([]rune, 0)
	if unicode.IsDigit(runes[0]) {
		return "", IncorrectStringError
	}

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			number := make([]rune, 0)
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				number = append(number, runes[i])
				i++
			}
			i--
			n, _ := strconv.Atoi(string(number))
			for j := 0; j < n-1; j++ {
				res = append(res, res[len(res)-1])
			}
		} else {
			res = append(res, runes[i])
		}
	}
	return string(res), nil
}

func main() {
	s, err := unpackString("ab5")
	if err != nil {
		return
	}
	fmt.Println(s)
}
