package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	c := make(chan requestResult)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.raddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}

}

var errRequestFailed = errors.New("request failed")

func hitUrl(url string, c chan requestResult) {

	fmt.Println("checking : ", url)

	resp, err := http.Get(url)

	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- requestResult{url: url, status: status}
}
