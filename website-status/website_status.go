package main

import (
	"fmt"
	"net/http"
	"sync"
)

func printStatusCode(url string) {
	res := getHTTPResponse(url)
	fmt.Printf("[%s] Code: %d\n", url, res.StatusCode)
}

func getHTTPResponse(url string) *http.Response {
	res, _ := http.Get(url)
	return res
}

func mai1n() {
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

