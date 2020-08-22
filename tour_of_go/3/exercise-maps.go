package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	a := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		_, ok := a[v]
		if ok {
			a[v]++
		} else {
			a[v] = 1
		}
	}
	return a
}

func main() {
	wc.Test(WordCount)
}
