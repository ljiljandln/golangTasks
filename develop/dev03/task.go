package main

import (
	"bufio"
	"bytes"
	"flag"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Flags struct {
	k       int
	n, r, u bool
	path    string
}

func NewFlags() Flags {
	path := flag.String("f", "", "имя файла")
	k := flag.Int("k", 1, "сортировка по колонке")
	n := flag.Bool("n", false, "сортировка по числовому значению")
	r := flag.Bool("r", false, "сортировка в обратном порядке")
	u := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()
	return Flags{k: *k, n: *n, r: *r, u: *u, path: *path}
}

type Data struct {
	keys []string
	m    map[string]int
}

func parseFile(path string) (Data, error) {
	in, err := os.Open(path)
	defer in.Close()
	if err != nil {
		return Data{}, err
	}

	m := make(map[string]int)
	var keys []string
	scanner := bufio.NewScanner(in)

	for scanner.Scan() {
		key := scanner.Text()
		if _, ok := m[key]; ok {
			m[key]++
		} else {
			m[key] = 1
			keys = append(keys, key)
		}
	}
	return Data{keys: keys, m: m}, nil
}

func sortN(keys []string) {
	sort.Slice(keys, func(i, j int) bool {
		vi, _ := strconv.Atoi(keys[i])
		vj, _ := strconv.Atoi(keys[j])
		return vi < vj
	})
}

func sortSimple(keys []string) {
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
}

func sortR(keys []string) {
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] > keys[j]
	})
}

func sortK(keys []string, k int) {
	var keysLessK, keysHasK []string

	for _, key := range keys {
		columns := strings.Fields(key)
		if len(columns) < k {
			keysLessK = append(keysLessK, key)
		} else {
			keysHasK = append(keysHasK, key)
		}
	}

	sortSimple(keysLessK)
	sort.Slice(keysHasK, func(i, j int) bool {
		columns1 := strings.Fields(keysHasK[i])
		columns2 := strings.Fields(keysHasK[j])
		return columns1[k-1] < columns2[k-1]
	})

	keysLessK = append(keysLessK, keysHasK...)
	copy(keys, keysLessK)
}

func writeFile(data *Data, u bool) {
	out, err := os.Create("out.txt")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	var buffer bytes.Buffer
	for _, key := range data.keys {
		if u {
			buffer.Write([]byte(key))
			buffer.WriteRune('\n')
		} else {
			for i := data.m[key]; i > 0; i-- {
				buffer.Write([]byte(key))
				buffer.WriteRune('\n')
			}
		}
	}

	_, err = out.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
}

func (data *Data) sorting(f *Flags) {
	if f.k != 1 {
		sortK(data.keys, f.k)
	} else if f.n {
		sortN(data.keys)
	} else if f.r {
		sortR(data.keys)
	} else {
		sortSimple(data.keys)
	}
}

func solve(data *Data, f *Flags) {
	data.sorting(f)
	writeFile(data, f.u)
}

func main() {
	f := NewFlags()
	data, err := parseFile(f.path)
	if err != nil {
		panic(err)
	}
	solve(&data, &f)
}
