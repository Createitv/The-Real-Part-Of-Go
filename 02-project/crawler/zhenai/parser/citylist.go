package parser

import (
	"The-Real-Part-Of-Go/02-project/crawler/engine"
	"fmt"
	"regexp"
)

//<a href="http://www.zhenai.com/zhenghun/xinbei" data-v-1573aa7c>新北</a>
const cityListRe = `<a href="(https?://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>(.*?)</a>`

func ParserCityList(content []byte) engine.ParseResult {
	//<a href="http://www.zhenai.com/zhenghun/jianghan" class="black" data-v-690ff80c>
	re, _ := regexp.Compile(cityListRe)
	matches := re.FindAllStringSubmatch(string(content), -1)
	result := engine.ParseResult{}
	for _, v := range matches {
		result.Items = append(result.Items, v[2])
		result.Requests = append(result.Requests, engine.Request{Url: string(v[1]), ParserFunc: ParserCityUser}) // 返会到CityList解析器
		fmt.Printf("URL:%s City:%s\n", string(v[1]), string(v[2]))
	}
	//fmt.Println(len(matches))
	return result
}
