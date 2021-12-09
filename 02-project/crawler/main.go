package main

import (
	"The-Real-Part-Of-Go/02-project/crawler/engine"
	"The-Real-Part-Of-Go/02-project/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParserCityList,
	})

	//str := `http:\u002F\u002Falbum.zhenai.com\u002Fu\u002F104237076`
	//fmt.Println(strconv.Unquote("\"" + str + "\""))
	//printCityList(all)
}
