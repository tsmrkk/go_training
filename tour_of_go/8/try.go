package main

import (
	"fmt"
	"time"
)

func main() {
	go say("hey")
	go say("there")
	time.Sleep(time.Millisecond * 600)
}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}
}
