package main

import (
	"fmt"
	"math"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Println(a, s)
}

func varialeInitialValue() {
	var a, b int = 3, 4
	var s string = "abd"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	a, b, c, g, h := 3, 4, 5, true, "s"
	b = 5
	fmt.Println(a, b, c, g, h)
}

func consts() {
	const fname = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(fname, c)
}

func enums() {
	const (
		cpp = iota
		java
		pythone
		golang
		javascript
	)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, java, pythone, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World")
	variableZeroValue()
	varialeInitialValue()
	variableTypeDeduction()
	consts()
	enums()
}

//DNS解析
//TCP连接
//发送HTTP请求
//服务器处理请求并返回HTTP报文
//浏览器解析渲染页面
//连接结束

// 浏览器向dns服务器解析域名获得ip，然后通过ip，向ip地址的服务器发起请求，然后服务器接收请求处理并返回http报文
