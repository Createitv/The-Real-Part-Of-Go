package parser

import (
	"The-Real-Part-Of-Go/02-project/crawler/engine"
	"The-Real-Part-Of-Go/02-project/crawler/model"
	"fmt"
	"regexp"
	"strconv"
)

//const ageRe = `<div .*?class="id">ID：(1751695861)</div>`
const ageRe = `<div .*?class="id">ID：(\d+)</div>`

//<div data-v-499ba28c="" class="des f-cl">武汉 | 26岁 | 大学本科 | 未婚 | 156cm | 5001-8000元<a data-v-499ba28c="" href="//www.zhenai.com/n/login?channelId=905819&amp;fromurl=https%3A%2F%2Falbum.zhenai.com%2Fu%2F1751695861" target="_self" class="online f-fr">查看最后登录时间</a></div>
const info = `<div.*?class="des f-cl">(.*?)<a.*?</div>`

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)
	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			// user age
			profile.Age = age
		}
	}

	infoRe := regexp.MustCompile(info)
	matchInfo := infoRe.FindSubmatch(contents)
	if match != nil {
		info := string(matchInfo[1])

		profile.Info = info

	}
	fmt.Printf("info", profile.Info, profile.Age)
	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}
