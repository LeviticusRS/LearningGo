package main

import (
	"fmt"
	"net/http"
	"sync"
)

func printStatusCode(url string) {
	res := getHttpResponse(url)
	fmt.Printf("[%s] Code: %d\n", url, res.StatusCode)
}

func getHttpResponse(url string) *http.Response {
	res, _ := http.Get(url)
	return res
}

func main() {
	var wg sync.WaitGroup
	urls := [...]string{
		"https://google.com",
		"https://facebook.com",
		"https://dev.to",
		// ... thousands more
	}
	for i, url := range urls {
		wg.Add(1)
		go func(pid int, url string) {
			printStatusCode(url)
			wg.Done()
		}(i, url)
	}
	wg.Wait()
}

