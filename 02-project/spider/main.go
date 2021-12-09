package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {

	//req, err := http.NewRequest("GET", "https://www.cnblogs.com/", nil)
	//req, err := http.NewRequest("GET", "https://book.douban.com/", nil)
	req, err := http.NewRequest("GET", "https://chinanews.com/", nil)
	if err != nil {
		panic(err)
	}
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}
	for key, value := range header {
		req.Header.Add(key, value)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: %v\n", resp.StatusCode)
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", result)
}
