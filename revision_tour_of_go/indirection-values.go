package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) tryChangeValue() float64 {
	v.X = 2
	v.Y = 3
	return 0
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	p.tryChangeValue()
	// methodの場合はp.abs()を(*p).abs()とよしなに変換してくれる
	// ただしp.tryChangeValue()だとp自体の値は変更できない
	fmt.Println(*p)
	fmt.Println(AbsFunc(*p))
}
