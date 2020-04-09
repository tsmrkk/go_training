package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1} //Yは0になる
	v3 = Vertex{}     //X, Yは0になる
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
