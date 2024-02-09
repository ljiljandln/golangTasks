package main

import (
	"strings"
	"testing"
)

func TestAnagram(t *testing.T) {
	actual := solve([]string{"тяпка", "пЯтак", "пятка", "клОун", "локун"})
	expected := make(map[string][]string)
	expected["тяпка"] = []string{"пятак", "пятка", "тяпка"}
	expected["клоун"] = []string{"клоун", "локун"}

	if len(expected) != len(actual) {
		t.Errorf("invalid size")
	}

	for key, expectedValue := range expected {
		actualValue, exist := actual[key]
		val := *actualValue
		if !exist {
			t.Errorf("the key %s is missing in actual", key)
		}
		if len(expectedValue) != len(val) {
			t.Error("the sizes of slices are not equal")
		}
		for i := 0; i < len(val); i++ {
			if strings.Compare(expectedValue[i], val[i]) != 0 {
				t.Errorf("the strings are different")
			}
		}
	}
}

func TestEmptyRes(t *testing.T) {
	res := solve([]string{"мышь", "кот"})
	if len(res) != 0 {
		t.Errorf("wrong result")
	}
}
