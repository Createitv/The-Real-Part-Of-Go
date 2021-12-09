package main

import (
	"fmt"
	"io/ioutil"
	"math"
)

// 内置对象类型bool, string, int int8, int16, int32, int64, uintptr（指针类型）
//  -- byte, rune(单字节，单字符,int32)
//  -- float, float32, float64, complex64, complex128 复数类型，数学计算使用
//  -- 强类型语言，非js隐式类型转化
//  -- 枚举类型enums
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variableTypeDeducation() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func varibleShorted() {
	a, b, c := 3, 4, true
	fmt.Println(a, b, c)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3, 4
	c := int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)

}

//枚举类型
func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, golang, python)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

// if语句
func readFile(filename string) {

	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s", contents)
	//}

	// if语句可以接受表达式赋值,if语句的作用域可以在else中使用
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}

}

//switch语句, 自动break， 除非fallthrough会自动向下执行,不然满足后直接退出
func eval(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operation" + op)

	}
	return result
}

// 循环
func doSum() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

// 指针,golang参数传递都是值传递，所谓的引用传递在go中传递的引用对象的地址值
func pointUse() {
	var a int = 2
	var pa *int = &a
	*pa = 3
	fmt.Println(a)
}

// 值传递
func referenceByValue() {
	fmt.Println("默认值传递")
	s := 100
	p := s // s 的地址给了p

	fmt.Println(p) //100 所以p对应的值变成了100
	p += 100
	fmt.Println(p, s) //200
	//200 p和s相同的地址，所以s也变成200
}

// 引用传递
func referenceByAddress() {
	fmt.Println("地址传递")
	s := 100
	p := &s // s 的地址给了p

	fmt.Println(*p) //100 所以p对应的值变成了100
	*p += 100
	fmt.Println(*p, s) //200
	//200 p和s相同的地址，所以s也变成200
}

// 交换变量
func swapTwoVariable(a, b *int) {
	fmt.Println(*a, *b)
	*a, *b = *b, *a
	fmt.Println(*a, *b)
}

// 内置容器： 数组初始化必须确定大小, 依然式值类型 []int代表切片 [10]int代表数组
func toArray() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 4}         // 初始化
	arr3 := [...]int{2, 4, 5, 6, 7} // 编译器数数组大小
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// i 索引序号 v值
	for i, v := range arr3 {
		fmt.Println(i, arr3[i], v)
	}
	for _, v := range arr3 {
		fmt.Println(v)
	}
}

// 内置容器：切片 基本和python列表无异 len(数组长度)cap(数组容量)
func toSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s := arr[2:6]
	s[1] = 100
	s = append(s, 111)
	fmt.Print(s, arr[:])
	fmt.Println(len(s), cap(s))
}

// slice动态扩容触发限制cap增加一倍
func silceAppend() {
	var s []int
	s2 := make([]int, 16)
	s3 := make([]int, 16, 32)
	fmt.Println(s2, cap(s3))
	for i := 0; i < 150; i++ {
		//fmt.Println(len(s), cap(s))
		s = append(s, 2*i+1)
	}
	fmt.Println(s)
}

// 内置容器：Map[K]V 哈希无序索引效率log1
func toMap() {
	m := map[string]string{
		"name":    "cool",
		"course":  "golang",
		"site":    "imooc",
		"quality": "high",
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
	course := m["course"]
	fmt.Println(course)
	cours := m["cours"]
	fmt.Println(cours)
	cour, ok := m["cours"]
	fmt.Println(cour, ok)
}

// go OOP面向对象比较特别 没有class 只有struct
// 结构体 go语言只支持封装，不支持多态和继承

func toStruct() {
	//type treeNode struct {
	//	value int
	//	left, right *treeNode
	//}
	var root treeNode
	fmt.Println(root) //{0 <nil> <nil>}
	root = treeNode{value: 3}

	/*
		root.left = &treeNode{}
		root.right = &treeNode{}
		fmt.Println(root)  //{3 0xc000004090 0xc0000040a8} 左右树都是内存地址
	*/
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	fmt.Println(root.right) //&{5 0xc0000040c0 <nil>}

	// 树的slice
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes) //[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc000004078}]

	root.left.right = creatNodeFactory(2)
	fmt.Println(root.left.right) // &{2 <nil> <nil>}

	root.print()
	root.left.right.print()

	for _, v := range nodes {
		fmt.Println(v)
	}
}

type treeNode struct {
	value       int
	left, right *treeNode
}

// 在结构体上定义方法,node相当于方法的接收者,依旧是值传递
func (node *treeNode) print() {
	fmt.Println(node.value)
}

// struct的工程函数
func creatNodeFactory(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	//fmt.Print("Hello World!")
	//variableZeroValue()
	//variableInitialValue()
	//variableTypeDeducation()
	//varibleShorted()
	//consts()
	//enums()
	//readFile("./test.txt")
	//fmt.Println(eval(3,4,"+"))
	//doSum()
	//pointUse()
	//referenceByValue()
	//referenceByAddress()
	//a, b := 3, 4
	//swapTwoVariable(&a, &b)
	//toArray()
	//toSlice()
	//silceAppend()
	//toMap()
	toStruct()
}
