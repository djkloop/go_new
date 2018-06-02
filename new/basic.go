package main

import (
	"fmt"
	"math"
)

var aInt = 3333
var bString = "bbbb"
var cBoolean = true

var (
	dInt     = 444
	eString  = "rrr"
	fBoolean = false
)

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableType() {
	a, b, c, s := 3, 4, true, "qqq"
	fmt.Println(a, b, c, s)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		filenames = "abcc.txt"
		cc        = 4
	)
	const filename string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
	fmt.Println(filenames, cc)
}

func enums() {
	const (
		cpp = iota
		java
		python
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

	fmt.Println(cpp, javascript, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)

}

func main() {
	fmt.Println("Hello World!")
	variable()
	variableInitialValue()
	variableTypeDeduction()
	variableType()
	fmt.Println(aInt, bString, cBoolean)
	fmt.Println(dInt, eString, fBoolean)
	triangle()
	consts()
	enums()
}
