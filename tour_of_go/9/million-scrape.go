package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	url := "http://localhost:8080"
	maxConnection := make(chan bool, 200)
	wg := &sync.WaitGroup{}
	count := 0
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		maxConnection <- true
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				return
			}
			defer resp.Body.Close()
			count++
			<-maxConnection
		}()
	}
	wg.Wait()
	end := time.Now()
	log.Printf("%d回のリクエストの送信に成功\n", count)
	log.Printf("%f秒かかりました\n", (end.Sub(start)).Seconds())
}
