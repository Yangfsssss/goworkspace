package main

import (
	"bufio"
	"byteCounter"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// var c byteCounter.ByteCounter
	// c.Write([]byte("hello"))
	// fmt.Println(c)
	// c = 0

	// var name = "Dolly"
	// fmt.Fprintf(&c, "Hello, %s", name)
	// fmt.Println(c)
	// countWords()

	// testCountingWriter()

	// testTreeValue()

	// testHTMLParser()

	testLimitReader()
}

func countWords() {
	var c byteCounter.ByteCounter
	c.CountWordsAndLines()
}

func testCountingWriter() {
	// bytes.Buffer 是 Go 标准库中的一个类型，它实现了 io.Writer 和 io.Reader 接口，并提供了对内存缓冲区的读写操作
	// 这种方式可以方便地创建和使用一个新的 bytes.Buffer，而无需显式地使用 new 关键字
	// buffer := &bytes.Buffer{}
	buffer := new(bytes.Buffer)
	writer, count := byteCounter.CountingWriter(buffer)

	fmt.Fprintf(writer, "Hello World!")
	fmt.Println("Bytes written", *count)

	fmt.Fprintf(writer, "9999")
	fmt.Println("Bytes written", *count)
}

func testTreeValue() {
	var t *byteCounter.Tree

	t = t.Insert(5)
	t = t.Insert(3)
	t = t.Insert(7)
	t = t.Insert(1)
	t = t.Insert(9)

	fmt.Println(t.String())
}

func testHTMLParser() {
	html := `<html>
		<div><h1>hello</h1></div>
	</html>`

	byteCounter.HTMLParser(html)
}

func testLimitReader() {
	// 数据源
	input := "Hello, World!"
	// 生成数据源的reader
	reader := strings.NewReader(input)

	// 将初始reader包装为LimitReader
	limitReader := byteCounter.LimitReader(reader, 5)
	// 生成储存的缓冲区
	buf := make([]byte, 10)
	// 读取
	_, err := limitReader.Read(buf)
	// fmt.Println(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	bufReader := bytes.NewReader(buf)
	// bytes 包中没有 NewScanner 函数。
	// 这是因为 bytes 包主要是提供了一些操作 []byte 类型的函数，而不是专门用于处理输入流的函数
	scanner := bufio.NewScanner(bufReader)
	// 指定每个字节作为分隔符
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// fmt.Println(string(buf[:n]))
}
