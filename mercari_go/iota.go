package main

import (
	"fmt"
)

func main() {
	const (
		a = iota
		b
	)
	const (
		c = 1 << iota
		d
		e
	)

	fmt.Println(a, b, c, d, e)
}
