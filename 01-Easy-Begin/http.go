package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func main() {
	res, err := http.Get("https://api.muxiaoguo.cn/api/dujitang")
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(res.Body)
	s, err := httputil.DumpResponse(res, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
