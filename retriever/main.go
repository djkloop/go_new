package main

import (
	"./mock"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string] string) string
}

const url = "http:www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post("http://www.baidu.com",
		map[string]string {
			"name": "djkloop",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string {
		"contents": "another faked baidu.com",
	})
	return s.Get(url)
}

func main() {
	reteriever := mock.Retriever{"good man!"}

	fmt.Println("Try a session")
	fmt.Println(session(&reteriever))
}


