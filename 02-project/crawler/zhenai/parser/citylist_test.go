package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	//contents, _ := fetcher.Fetch("https://www.zhenai.com/zhenghun")
	// 读取本地爬取的数据
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParserCityList(contents)
	const resultSize = 470

	expectedCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests, got %d", resultSize, len(result.Items))
	}
	//fmt.Printf("%s\n", contents)
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, got %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expectedUrl #%d: %s; but got %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expectedUrl #%d: %s; but got %s", i, city, result.Items[i].(string))
		}
	}

	//fmt.Printf("%s\n", contents)
}
