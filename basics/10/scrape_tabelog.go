package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("https://tabelog.com/tokyo/A1305/")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	// doc, err := html.Parse(byteArray)
}
