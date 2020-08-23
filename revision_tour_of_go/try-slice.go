package main

import "fmt"

func main() {
	s := [6]int{2, 3, 5, 7, 11, 13}
	var a []int

	a = s[:2]
	printSlice(a)

	a = s[:4]
	printSlice(a)

	a = s[2:]
	printSlice(a)

	a = append(a, 100)

	printSlice(a)

	a[0] = 10
	a[4] = 900

	printSlice(a)

	fmt.Println("##############################")

	fmt.Println("s is ", s)
	fmt.Println("##############################")
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
