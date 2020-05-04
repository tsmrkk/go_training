package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	url := "https://qiita.com/Azunyan1111/items/a1b6c58dc868814efb51"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
	tokenizer := html.NewTokenizer(resp.Body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			return
		}
		if tokenType == html.StartTagToken {
			token := tokenizer.Token()
			if "title" == token.Data {
				tokenType = tokenizer.Next()
				if tokenType == html.TextToken {
					fmt.Println(tokenizer.Token().Data)
					break
				}
			}
		}
	}
}
