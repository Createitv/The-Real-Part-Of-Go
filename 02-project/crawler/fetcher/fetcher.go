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
)

func Fetch(url string) ([]byte, error) {
	//resp, err := http.Get("https://www.zhenai.com/zhenghun")
	//resp, err := http.Get("https://www.zhenai.com/zhenghun/wuhan")
	//resp, err := http.Get("https://pvp.qq.com/")
	resp, err := http.Get(url)
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
	e := determineEncoding(resp.Body)
	fmt.Print(e)
	utf8Reader := transform.NewReader(resp.Body, e.NewEncoder())
	return ioutil.ReadAll(utf8Reader)

}
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)
	if err != nil {
		log.Printf("Fetch error: %v", err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
