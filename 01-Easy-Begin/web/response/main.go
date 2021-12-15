package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func responseBody(r *http.Response) {
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s", content)
}

func status(r *http.Response) {
	fmt.Println(r.StatusCode)
	fmt.Println(r.Status)
}

func header(r *http.Response) {
	fmt.Println(r.Header.Get("content-encoding"))
	fmt.Println(r.Header.Get("Host"))
	fmt.Println(r.Header.Get("content-type"))
	fmt.Println(r.Header.Get("user-agent"))
	fmt.Println(r.Header.Values("args"))
}

func encoding(r *http.Response) {
	buf := bufio.NewReader(r.Body)
	bytes, _ := buf.Peek(1024)
	e, _, _ := charset.DetermineEncoding(bytes, r.Header.Get("content-type"))
	bodyReader := transform.NewReader(buf, e.NewDecoder())
	content, _ := ioutil.ReadAll(bodyReader)
	fmt.Printf("%s", content)

}

func main() {
	request, err := http.NewRequest(http.MethodGet, "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	//params := make(url.Values)
	params := url.Values{}
	params.Add("name", "talkBeCheap")
	params.Add("age", "18")
	params.Add("User-Agent", "Mozilla5.0")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
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
	responseBody(r)
	status(r)
	header(r)
	encoding(r)
}
