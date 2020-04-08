package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const ooo = 3.2
	const Truth = true
	fmt.Println("Go rules?", Truth)
	fmt.Println("Go rules?", ooo)
}
