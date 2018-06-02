package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d", score))
	case score < 60:
		g = "没有及格"
	case score < 80:
		g = "一般般啦"
	case score < 90:
		g = "继续努力"
	case score <= 100:
		g = "骚的一p"
	}
	return g
}

func main() {
	const filename = "./new/abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(grade(90))
}
