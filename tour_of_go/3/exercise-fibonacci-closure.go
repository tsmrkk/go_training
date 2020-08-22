package main

import "fmt"

func fibonacci() func(int) int {
	a := 0
	b := 1
	return func(x int) int {
		b = a + b
		a = b - a
		return b
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
