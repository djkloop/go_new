package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Yes 双击评论666！"
	fmt.Println(len(s))

	for i, char := range []rune(s) {
		fmt.Printf("(%d %c)", i, char)
	}
	fmt.Println()

	ss := "hidfdf"
	fmt.Println(strings.Fields(ss))
}
