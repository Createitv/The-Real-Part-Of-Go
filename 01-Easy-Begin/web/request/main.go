package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printBody(r *http.Response) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("content: %s\n", string(content))
}

func requestByParams() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	//params := make(url.Values)
	params := url.Values{}
	params.Add("name", "talkBeCheap")
	params.Add("age", "18")
	fmt.Println(params.Encode())
	request.URL.RawQuery = params.Encode()

	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	printBody(r)
}

func requestByHeaders() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("user-Agent", "Chrome5.0 Mac OS X")
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	printBody(r)
}

func main() {
	// 设置请求查询参数
	//requestByParams()
	// 如何定制请求头
	requestByHeaders()
}
