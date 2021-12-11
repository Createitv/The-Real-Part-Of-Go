package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	//https://album.zhenai.com/u/1164458612
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authority", "www.zhenai.com")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Sec-Ch-Ua", "^^")

	resp, err := client.Do(req)

	//resp, err := http.Get("https://www.zhenai.com/zhenghun")
	//resp, err := http.Get("https://www.zhenai.com/zhenghun/wuhan")
	//resp, err := http.Get("https://pvp.qq.com/")
	//resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong StatusCode:%d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	fmt.Print(e)
	utf8Reader := transform.NewReader(bodyReader, e.NewEncoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetch error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
