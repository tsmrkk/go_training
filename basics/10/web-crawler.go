package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	crawl(url, depth, fetcher, ch)
}

func isInclude(arr []string, url string) bool {
	for _, v := range arr {
		if v == url {
			return true
		}
	}
	return false
}

func crawl(url string, depth int, fetcher Fetcher, ch chan string) {
	if depth <= 0 {
		return
	}
	if isInclude(arr, url) {
		return
	}
	fmt.Println("##################################")
	ch <- url
	// fmt.Println("ch", len(ch))
	fmt.Println("##################################")
	arr = append(arr, url)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	var wg sync.WaitGroup
	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			crawl(url, depth-1, fetcher, ch)
			defer wg.Done()
		}(u)
	}
	wg.Wait()
	return
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

var arr []string

func main() {
	ch := make(chan string)
	Crawl("https://golang.org/", 4, fetcher, ch)
}
