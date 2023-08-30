package byteCounter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
func (c *ByteCounter) CountWordsAndLines() (wordsCount int, linesCount int) {
	content, err := os.ReadFile(os.Args[1:][0])
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(content))

	for scanner.Scan() {
		// scanner.Text() 不能直接返回一行文本内容，而是返回当前扫描器位置到上一个分隔符之间的文本片段
		// 在默认情况下，bufio.Scanner 使用换行符作为分隔符

		// bufio.Scanner 默认使用 bufio.ScanLines 作为分割函数
		// 如果只是想获取文本的行数，那么使用 bufio.ScanLines 更适合。

		// 获取一行的文本
		line := scanner.Text()
		fmt.Println("line:", line)
		// 将一行的文本包装成获取行中单词的scanner
		words := bufio.NewScanner(strings.NewReader(line))
		// 设置scanner的split函数
		words.Split(bufio.ScanWords)

		for words.Scan() {
			// fmt.Printf("%s word:", words.Text())
			wordsCount++
		}

		linesCount++
	}

	// fmt.Println(content)

	// advance, token, err := bufio.ScanWords(content, false)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%d words, %d lines\n", wordsCount, linesCount)
	return
}

// 练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
type countingWriter struct {
	w     io.Writer
	count *int64
}

func (cw *countingWriter) Write(p []byte) (int, error) {
	n, err := cw.w.Write(p)
	*cw.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	count := int64(0)
	return &countingWriter{w, &count}, &count
}

// 练习 7.3： 为在gopl.io/ch4/treesort（§4.4）中的*tree类型实现一个String方法去展示tree类型的值序列。
type Tree struct {
	value int
	right *Tree
	left  *Tree
}

// strings.Builder 是 Go 标准库中的一个类型，用于高效地构建字符串。
// 它提供了一些方法来追加字符串、字节和字符到内部的缓冲区，并可以通过调用其 String() 方法将缓冲区中的内容作为字符串返回。
func (t *Tree) stringBuilder(builder *strings.Builder) {
	if t == nil {
		return
	}

	t.left.stringBuilder(builder)
	// Sprintf formats according to a format specifier and returns the resulting string.
	builder.WriteString(fmt.Sprintf("%d ", t.value))
	t.right.stringBuilder(builder)
}

func (t *Tree) String() string {
	var builder strings.Builder
	t.stringBuilder(&builder)
	return builder.String()
}

func (t *Tree) Insert(value int) *Tree {
	if t == nil {
		return &Tree{value: value}
	}

	if value < t.value {
		t.left = t.left.Insert(value)
	} else {
		t.right = t.right.Insert(value)
	}

	return t
}

// 练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。
// 实现一个简单版本的NewReader，用它来构造一个接收字符串输入的HTML解析器（§5.2）
func simpleNewReader(s string) io.Reader {
	return strings.NewReader(s)
}

func HTMLParser(s string) {
	// 生成reader
	reader := simpleNewReader(s)
	content, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}

	var buf strings.Builder
	// html.Render 函数会将 *html.Node 的内容以字符串的形式写入到 strings.Builder 中
	html.Render(&buf, content)

	fmt.Println(buf.String())
}

// 练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数：
type limitReader struct {
	r io.Reader
	n int64
}

func (lr *limitReader) Read(p []byte) (n int, err error) {
	// 自身没有长度，返回
	if lr.n <= 0 {
		return 0, io.EOF
	}

	// 读取长度大于自身长度，将读取长度更改为自身长度
	if int64(len(p)) > lr.n {
		p = p[0:lr.n]
	}

	// 方法 Read，用于从数据源中读取数据并将其存储到指定的缓冲区中
	// 读取
	n, err = lr.r.Read(p)
	// 减去已读取长度，即更新自身可读取长度
	lr.n -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}
