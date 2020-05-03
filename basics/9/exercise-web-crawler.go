package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Checked struct {
	urls []string
	mux  sync.Mutex
}

var wg sync.WaitGroup

func Crawl(url string, depth int, fetcher Fetcher, ch Checked) {
	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	ch.urls = append(ch.urls, url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		flag := -1
		for _, v := range ch.urls {
			if v == u {
				flag = 1
			}
		}
		if flag == -1 {
			wg.Add(1)
			ch.mux.Lock()
			go func() {
				Crawl(u, depth-1, fetcher, ch)
				defer wg.Done()
				defer ch.mux.Unlock()
			}()
		}
		wg.Wait()
	}
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

func main() {
	ch := Checked{}
	Crawl("https://golang.org/", 4, fetcher, ch)
}
