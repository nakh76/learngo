package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	var results = make(map[string]string)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.raddit.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
	}

	results["hello"] = "hi"

	for _, url := range urls {
		result := "OK"
		err := hitUrl(url)
		if err != nil {
			result = "FAIL"
		}
		results[url] = result
	}

	for url, result := range results {
		fmt.Println(url, result)
	}

}

var errRequestFailed = errors.New("request failed")

func hitUrl(url string) error {

	fmt.Println("checking : ", url)

	resp, err := http.Get(url)

	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}

	return nil
}
