package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

func main() {
	// 一个类型如果拥有一个接口需要的所有方法，那么这个类型就实现了这个接口
	var w io.Writer

	w = os.Stdout
	w = new(bytes.Buffer)
	// w = time.Second // compile error: time.Duration lacks Write method

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

	w = rwc
	// rwc = w // compile error: io.Writer lacks Close method

	// 在T类型的参数上调用一个*T的方法是合法的，只要这个参数是一个变量；
	// 编译器隐式的获取了它的地址。但这仅仅是一个语法糖：T类型的值不拥有所有*T指针的方法
	type IntSet struct{}
	// func (*IntSet) String() {}

	// 接口类型封装和隐藏具体类型和它的值。即使具体类型有其它的方法，也只有接口类型暴露出来的方法会被调用到
	os.Stdout.Write([]byte("hello"))
	os.Stdout.Close()

	var ww io.Writer
	ww = os.Stdout
	ww.Write([]byte("hello"))
	// w.Close() // compile error: io.Writer lacks Close method

	// 因为空接口类型对实现它的类型没有要求，所以我们可以将任意一个值赋给空接口类型
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)

	// Album
	// Book
	// Movie
	// Magazine
	// Podcast
	// TVEpisode
	// Track

	type Artifact interface {
		Title() string
		Creators() string
		Created() time.Time
	}

	type Text interface {
		Pages() int
		Words() int
		PageSize() int
	}

	type Audio interface {
		Stream() (io.ReadCloser, error)
		RunningTime() time.Duration
		Format() string
	}

	type Video interface {
		Stream() (io.ReadCloser, error)
		RunningTime() time.Duration
		Format() string
		Resolution() (x, y int)
	}

	type Streamer interface {
		Stream() (io.ReadCloser, error)
		RunningTime() time.Duration
		Format() string
	}

}
