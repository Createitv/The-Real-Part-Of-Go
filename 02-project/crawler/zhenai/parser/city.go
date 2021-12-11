package parser

import (
	"The-Real-Part-Of-Go/02-project/crawler/engine"
	"fmt"
	"regexp"
)

const cityRex = `<a href="(https?://album.zhenai.com/u/[0-9]+)".*? alt="(.*?)"></a>`

//const cityListRe = `<a href="(https?://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>(.*?)</a>`

func ParserCityUser(content []byte) engine.ParseResult {
	re, _ := regexp.Compile(cityRex)
	matches := re.FindAllStringSubmatch(string(content), -1)
	result := engine.ParseResult{}
	for _, v := range matches {
		result.Items = append(result.Items, v[2])
		result.Requests = append(result.Requests, engine.Request{Url: string(v[1]), ParserFunc: ParseProfile})
		fmt.Printf("CityURL:%s City:%s\n", string(v[1]), string(v[2]))
	}
	//fmt.Println(len(matches))
	return result
}
