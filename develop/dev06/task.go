package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type flags struct {
	f int
	d string
	s bool
}

func getFlags() (fl flags) {
	flag.IntVar(&fl.f, "f", 1, "выбрать поля (колонки)")
	flag.StringVar(&fl.d, "d", " ", "использовать другой разделитель (дефолтно: space)")
	flag.BoolVar(&fl.s, "s", false, "только строки с разделителем")
	flag.Parse()
	return
}

func getLines() (lines []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return
}

func makeCut(fl *flags, lines []string) (res []string) {
	for _, line := range lines {
		spl := strings.Split(line, fl.d)
		if fl.s && !strings.Contains(line, fl.d) {
			continue
		} else if len(spl) >= fl.f {
			res = append(res, spl[fl.f-1])
		} else {
			res = append(res, line)
		}
	}
	return
}

func printRes(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func main() {
	fl := getFlags()
	lines := getLines()
	res := makeCut(&fl, lines)
	printRes(res)
}
