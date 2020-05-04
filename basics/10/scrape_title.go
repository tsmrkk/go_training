package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := "https://qiita.com/Azunyan1111/items/a1b6c58dc868814efb51"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	doc := html.NewTokenizer(strings.NewReader(string(byteArray)))
	for {
		tt := doc.Next()
		if tt == html.ErrorToken {
			return
		}
		// fmt.Println(tt.String())
		tn, _ := doc.TagName()
		if len(tn) == 1 && tn[0] == "title" {
			fmt.Println(tt)
			break
		}
	}
	// fmt.Println(string(byteArray))

	// var f func(*html.Node)
	// f = func(n *html.Node) {
	// 	if n.Type == html.ElementNode && n.Data == "title" {
	// 		fmt.Println(n.Token())
	// 	}
	// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
	// 		f(c)
	// 	}
	// }

	// f(doc)
}
