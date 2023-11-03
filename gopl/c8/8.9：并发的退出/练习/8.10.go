// 练习 8.10： HTTP请求可能会因http.Request结构体中Cancel channel的关闭而取消。
// 修改8.6节中的web crawler来支持取消http请求。
// （提示：http.Get并没有提供方便地定制一个请求的方法。你可以用http.NewRequest来取而代之，
// 设置它的Cancel字段，然后用http.DefaultClient.Do(req)来进行这个http请求。）

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"

	"golang.org/x/net/html"
)

type Page struct {
	URL   string
	Depth int
}

var (
	done        = make(chan struct{})
	ctx, cancel = context.WithCancel(context.Background())
)

func main() {
	runCrawl(1)
}

func runCrawl(depthLimit int) {
	workList := make(chan Page)
	seen := make(map[string]bool)

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
		workList <- Page{os.Args[1], 0}
	}()

	go func() {
		wg.Wait()
		close(workList)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("cancel signal received,program shutting down-------------------------------------------------------------------")
		cancel()
		close(done)
	}()

	// 要在特定条件下立即退出循环，可以使用 break 语句来实现。
	// 如果需要在循环之外的地方控制循环的退出，那么使用 break loop 可能更加方便。
	// for i := 0; i < 20; i++ {
loop:
	for {
		select {
		case <-done:
			fmt.Println("main goroutine closed")
			return
		case page, ok := <-workList:
			if !ok {
				break loop
			}
			wg.Add(1)
			go func() {
				defer wg.Done()
				// for page := range workList {
				if page.Depth > depthLimit {
					// break loop
					return
				}

				fmt.Println("Crawling:", page.URL)
				links := crawl(page.URL)
				fmt.Println("Crawling Links", links)

				mu.Lock()
				for _, link := range links {
					if !seen[link] {
						seen[link] = true
						workList <- Page{link, page.Depth + 1}
					}
				}
				mu.Unlock()
				// }
			}()
		}
	}

	// }

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
	if canceled() {
		return nil
	}

	tokens <- struct{}{} // acquire a token
	list, err := Extract(url, ctx)
	<-tokens // release the token
	if err != nil {
		fmt.Println(err)
	}

	return list
}

func Extract(url string, ctx context.Context) ([]string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	select {
	case <-done:
		cancel()
		return nil, nil
	default:

		if resp.StatusCode != http.StatusOK {
			return nil, fmt.Errorf("getting %s", resp.Status)
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
		}

		var links []string
		visitNode := func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				for _, a := range n.Attr {
					if a.Key != "href" {
						continue
					}

					link, err := resp.Request.URL.Parse(a.Val)
					if err != nil {
						continue
					}

					links = append(links, link.String())
				}
			}
		}

		forEachNode(doc, visitNode, nil)
		return links, nil
	}

}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
