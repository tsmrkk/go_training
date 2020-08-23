package main

import (
	"fmt"
)

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// i.show()
}

func describe(i I) {
	fmt.Println(i)
}

// interfaceに対してのfunctionは作ることができるけど、methodは作れない

// func (i I) show() {
// 	fmt.Println(i)
// }
