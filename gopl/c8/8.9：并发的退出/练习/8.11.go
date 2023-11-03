package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	result      = make(chan []byte)
	done        = make(chan bool)
	ctx, cancel = context.WithCancel(context.Background())
)

func main() {
	urlList := make(chan string, 6)
	urlList <- "http://localhost:8899/data/1"
	urlList <- "http://localhost:8899/data/2"
	urlList <- "http://localhost:8899/data/3"
	urlList <- "http://localhost:8899/data/4"
	urlList <- "http://localhost:8899/data/5"
	urlList <- "http://localhost:8899/data/6"
	raceFetch(urlList)
	close(urlList)
}

func raceFetch(urlList chan string) {
	for {
		select {
		case value := <-result:
			fmt.Println("main goroutine closed,received value is: ", value)
			close(result)
			close(done)
			return
		default:
			url, ok := <-urlList
			if !ok {
				fmt.Println("error")
				return
			}
			go fetch(url, ctx)
		}
	}
}

func fetch(url string, ctx context.Context) {
	if canceled() {
		return
	}

	select {
	case <-done:
		fmt.Println("first value received,request shutting down-------------------------------------------------------------------")
		cancel()
		return
	default:
		req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
		// if err != nil {
		// 	return nil, err
		// }

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		}
		defer resp.Body.Close()

		// if resp.StatusCode != http.StatusOK {
		// 	return nil, fmt.Errorf("getting %s", resp.Status)
		// }

		res, _ := io.ReadAll(resp.Body)
		// if err != nil {
		// 	return nil, fmt.Errorf("ReadAll %s as resp.Body: %v", url, err)
		// }

		result <- res
		// done <- true
	}
}

func canceled() bool {
	select {
	case <-result:
		return true
	default:
		return false
	}
}
