package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFaild = errors.New("Request failed")

func main() {
	// 초기화 필요!!
	// var results map[string]string
	// var results = map[string]string{}
	// make를 활용하여 초기화 할수 있음
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}
	// 초기화 되지 않은 map에 추가하면 panic이 발생할수도 있음
	// results["hello"] = "Hello"
	// URL: OK | FAILED
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "FAILED"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errRequestFaild
	}
	return nil
}
