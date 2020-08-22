package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

const RANKING_SHOW_LIMIT int = 3

type record struct {
	val int
	key string
}

type result []record

func (rs result) Len() int { return len(rs) }

func (rs result) Less(i, j int) bool { return rs[i].val < rs[j].val }

func (rs result) Swap(i, j int) { rs[i], rs[j] = rs[j], rs[i] }

func removeIgnoreStrings(s string) string {
	ignoreStringArray := []string{".", ","}
	dls := s

	for _, is := range ignoreStringArray {
		dls = strings.ReplaceAll(dls, is, "")
	}

	return dls
}

func newDelimitedArray(s string) []string {
	delimiterStringArray := []string{" ", "\n"}
	for i := 1; i < len(delimiterStringArray); i++ {
		s = strings.ReplaceAll(s, delimiterStringArray[i], delimiterStringArray[0])
	}

	arr := strings.Split(s, delimiterStringArray[0])
	return arr
}

func newWordCountMap(arr []string) map[string]int {
	m := make(map[string]int)

	for _, v := range arr {
		_, ok := m[v]
		if ok {
			m[v]++
			continue
		}
		if v != "" {
			m[v] = 1
		}
	}

	return m
}

func output(ar []string) {
	for i := 0; i < RANKING_SHOW_LIMIT; i++ {
		fmt.Println(ar[i])
	}
}

func sortAlphabetically(rs result, l, r int) []string {
	var s []string
	for i := l; i < r; i++ {
		s = append(s, rs[i].key)
	}
	sort.Strings(s)
	return s
}

func newNumericallySortedResult(m map[string]int) result {
	var rs result

	for k, v := range m {
		rs = append(rs, record{key: k, val: v})
	}

	sort.Sort(sort.Reverse(rs))

	return rs
}

func newResultArray(rs result) []string {
	var rrs []string
	i := 0

	for i < RANKING_SHOW_LIMIT {
		j := i + 1
		for j < RANKING_SHOW_LIMIT && rs[i].val == rs[j].val {
			j++
		}
		if j != i+1 {
			s := sortAlphabetically(rs, i, j)
			for k := 0; k < len(s); k++ {
				rrs = append(rrs, s[k])
			}
			i = j
			continue
		}
		rrs = append(rrs, rs[i].key)
		i++
	}

	return rrs
}

func newInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return text
}

func main() {
	input := newInput()
	str := removeIgnoreStrings(input)
	arr := newDelimitedArray(str)
	m := newWordCountMap(arr)
	rs := newNumericallySortedResult(m)
	s := newResultArray(rs)
	output(s)
}
