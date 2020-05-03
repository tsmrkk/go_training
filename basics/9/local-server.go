package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
	return
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server started")
	http.ListenAndServe(":8080", nil)
}
