package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"testing"
)

const (
	input     = "in.txt"
	actualOut = "out.txt"
)

func TestTask(t *testing.T) {
	tests := []struct {
		name        string
		k           int
		n, r, u     bool
		expectedOut string
	}{
		{name: "simpleSortWithoutFlags", k: 1, n: false, r: false, u: false, expectedOut: "test_out/simple.txt"},
		{name: "sortWithFlagK", k: 2, n: false, r: false, u: false, expectedOut: "test_out/k.txt"},
		{name: "sortWithFlagN", k: 1, n: true, r: false, u: false, expectedOut: "test_out/n.txt"},
		{name: "sortWithFlagR", k: 1, n: false, r: true, u: false, expectedOut: "test_out/r.txt"},
		{name: "sortWithFlagU", k: 1, n: false, r: false, u: true, expectedOut: "test_out/u.txt"},
	}

	for _, test := range tests {
		data, _ := parseFile(input)
		f := Flags{test.k, test.n, test.r, test.u, input}
		solve(&data, &f)
		if getMD5(actualOut) != getMD5(test.expectedOut) {
			t.Errorf("Hashes are not equal on test %s\n", test.name)
		}
	}
}

func getMD5(path string) string {
	sum := md5.New()
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("wrong")
	}
	defer f.Close()
	io.Copy(sum, f)
	return fmt.Sprintf("%X", sum.Sum(nil))
}
