package main

import (
	"fmt"
	"net/http"
	"recur"

	"golang.org/x/net/html"
)

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// 虽然Go的垃圾回收机制会回收不被使用的内存，但是这不包括操作系统层面的资源，比如打开的文件、网络连接。
	// 因此我们必须显式的释放这些资源。
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	return recur.Visit(nil, doc), nil
}

// bare return
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// doc, err := html.Parse(resp.Body)
	_, err = html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing %s as HTML: %v", url, err)
		return
	}

	// to be implemented
	// words,images = countWordsAndImages(doc)
	return
}

func main() {}
