package main

import (
	"fmt"
	"links"
	"os"
	"sync"
)

// 练习 8.6： 为并发爬虫增加深度限制。也就是说，如果用户设置了depth=3，那么只有从首页跳转三次以内能够跳到的页面才能被抓取到。
type Page struct {
	URL   string
	Depth int
}

func runCrawl2(depthLimit int) {
	workList := make(chan Page)
	seen := make(map[string]bool)

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		workList <- Page{os.Args[1], 0}
	}()

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for page := range workList {
				if page.Depth > depthLimit {
					continue
				}

				fmt.Println("Crawling:", page.URL)
				links := crawl(page.URL)

				mu.Lock()
				for _, link := range links {
					if !seen[link] {
						seen[link] = true
						workList <- Page{link, page.Depth + 1}
					}
				}
				mu.Unlock()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(workList)
	}()

	for page := range workList {
		if page.Depth > depthLimit {
			break
		}
		fmt.Println(page.URL)
	}
}

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	// fmt.Println(url)

	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens // release the token
	if err != nil {
		fmt.Println(err)
	}

	return list
}
