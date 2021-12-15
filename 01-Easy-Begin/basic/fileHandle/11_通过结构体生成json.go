package main

import (
	"encoding/json"
	"fmt"
)

/*

{
    "company": "itcast",
    "subjects": [
        "Go",
        "C++",
        "Python",
        "Test"
    ],
    "isok": true,
    "price": 666.666
}

*/

//成员变量名首字母必须大写
//type IT struct {
//	Company  string
//	Subjects []string
//	IsOk     bool
//	Price    float64
//}

type IT struct {
	Company  string   `json:"-"`        //此字段不会输出到屏幕
	Subjects []string `json:"subjects"` //二次编码
	IsOk     bool     `json:",string"`
	Price    float64  `json:",string"`
}

func main() {
	//定义一个结构体变量，同时初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "Test"}, true, 666.666}
	s2 := IT{"MIT", []string{"golang", "c++"}, false, 422.2}
	//编码，根据内容生成json文本
	//{"Company":"itcast","Subjects":["Go","C++","Python","Test"],"IsOk":true,"Price":666.666}
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", "	") //格式化编码
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	j, _ := json.Marshal(s2)
	fmt.Println("buf = ", string(buf))
	fmt.Println(string(j))
	fmt.Println(len(j))
	for _, v := range j {
		fmt.Printf("k:%v k:%s\n", v, string(v))
	}
	//[123 34 115 117 98 106 101 99 116 115 34 58 91 34 103 111 108 97 110 103 34 44 34 99 43 43 34 93 44 34 73 115 79 107 34 58 34 102 97 108 115 101 34 44 34 80 114 105 99 101 34 58 34 52 50 50 46 50 34 125]
}
