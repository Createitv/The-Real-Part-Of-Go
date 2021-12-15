package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func Get() {
	r, err := http.Get("http://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("content: %s\n", string(content))
}

func Post() {
	r, err := http.Post("http://httpbin.org/post", "", nil)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(r.Body)
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("content: %s\n", string(content))
}

func Put() {
	request, err := http.NewRequest(http.MethodPut, "http://httpbin.org/put", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("content: %s\n", string(content))

}
func Del() {
	request, err := http.NewRequest(http.MethodDelete, "http://httpbin.org/delete", nil)
	if err != nil {
		panic(err)
	}
	r, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("content: %s\n", string(content))

}
func main() {
	// http基本方法，Get Post Put Del
	Get()
	Post()
	Put()
	Del()
}
