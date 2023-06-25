package main

import (
	"fmt"
	mock "mockretriever"
	"real"
	"time"
)

// 接口实现是隐式的
// 只要实现接口里的方法

// 接口变量自带指针
// 接口变量使用值传递，几乎不需要使用接口的指针
// 指针接收者实现只能以指针方式使用；值接收者都可

// 表示任何类型：interface{}
// Type Assertion
// Type Switch

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) string {
	return poster.Post(url, map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake imooc.com",
	})

	return s.Get(url)
}

func main() {
	var r Retriever

	retriever := mock.Retriever{Contents: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	fmt.Println("Try a session")
	fmt.Println(session(&retriever))

	// fmt.Println(download(r))

}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)
	fmt.Print(" > Type switch: ")

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	fmt.Println()
}
