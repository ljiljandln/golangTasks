package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	ReadingFileError = errors.New("can't read file")
	ExpressionError  = errors.New("can't get expression")
)

type Flags struct {
	path, pattern                 string
	after, before, context        int
	count, ignore, invert, number bool
}

func NewFlags() Flags {
	path := flag.String("f", "", "файл")
	pattern := flag.String("e", "", "паттерн")
	after := flag.Int("A", 0, "вывод N строк после совпадения")
	before := flag.Int("B", 0, "вывод N строк до совпадения")
	context := flag.Int("C", 0, "вывод N строк в районе совпадения")
	count := flag.Bool("c", false, "вывести количество строк с совпадением")
	ignore := flag.Bool("i", false, "игнорировать различия регистра")
	invert := flag.Bool("v", false, "инвертировать вывод")
	number := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()
	return Flags{*path, *pattern, *after, *before, *context, *count, *ignore, *invert, *number}
}

func readFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data, nil
}

func getExpression(pattern string, ignore bool) (*regexp.Regexp, error) {
	ignorePrefix := ""
	if ignore {
		ignorePrefix = "(?i)"
	}
	compiledExpression, err := regexp.Compile(ignorePrefix + pattern)
	if err != nil {
		return nil, err
	}
	return compiledExpression, nil
}

func getIntersectionsCount(file []string, expression *regexp.Regexp) (res int) {
	for _, str := range file {
		match := expression.Match([]byte(str))
		if match {
			res++
		}
	}
	return res
}

func search(file []string, expression *regexp.Regexp, fl *Flags) (res []string) {
	for i, str := range file {
		match := expression.Match([]byte(str))
		if (fl.invert && !match) || (!fl.invert && match) {
			res = echo(file, i, fl)
		}
	}
	return
}

func echo(file []string, i int, fl *Flags) (res []string) {
	startPoint := 0
	endPoint := len(file)
	if i+fl.after < len(file) {
		endPoint = i + fl.after
	}
	if i-fl.before >= 0 {
		startPoint = i - fl.before
	}
	if endPoint != len(file) {
		endPoint += 1
	}
	for line := startPoint; line < endPoint; line++ {
		prefix := ""
		if fl.number {
			prefix = fmt.Sprintf("%d: ", line+1)
		}
		str := fmt.Sprintf("%s%s", prefix, file[line])
		fmt.Println(str)
		res = append(res, str)
	}
	return
}

func solver(fl *Flags) ([]string, error) {
	file, err := readFile(fl.path)
	if err != nil {
		return []string{}, ReadingFileError
	}
	expression, err := getExpression(fl.pattern, fl.ignore)
	if err != nil {
		return []string{}, ExpressionError
	}

	var res []string
	if fl.count {
		count := getIntersectionsCount(file, expression)
		if fl.invert {
			count = len(file) - count
		}
		fmt.Println(count)
		res = append(res, strconv.Itoa(count))
	} else {
		if fl.after == 0 && fl.before == 0 && fl.context != 0 {
			fl.after = fl.context / 2
			fl.before = fl.context / 2
		}
		res = search(file, expression, fl)
	}
	return res, nil
}

func main() {
	fl := NewFlags()

	_, err := solver(&fl)
	if err != nil {
		fmt.Println(err)
		return
	}
}
