package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Fetch(url string) {
	req, _ := http.NewRequest("GET", url, nil)
	fmt.Println(url)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Sec-Ch-Ua", "\"Chromium\";v=\"94\", \"Google Chrome\";v=\"94\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")
	req.Header.Set("Sec-Ch-Ua-Platform", "\"macOS\"")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://spa1.scrape.center/")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	req.Header.Set("Cookie", "UM_distinctid=17d88f0495945a-09ea098800af2-123b6650-13c680-17d88f0495a76")

	defer req.Body.Close()
	res, _ := http.DefaultClient.Do(req)
	if res.StatusCode == http.StatusOK { // OK
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println(string(bodyBytes))
	}
	// return res, err
}

func main() {
	const offset = 10
	for i := 0; i < offset; i++ {
		url := "https://spa1.scrape.center/api/movie?limit=10&offset=" + strconv.Itoa(i*10)
		fmt.Println(url)
		go Fetch(url)
	}
	time.Sleep(time.Second * 2)

}
