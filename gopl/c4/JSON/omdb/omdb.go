package omdb

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const apiKey = "49e1e629"

const baseUrl = "http://www.omdbapi.com/" + "?apikey=" + apiKey

const dir = "/Users/bytedance/Documents/code/goworkspace/gopl/c4/JSON/omdb/posters"

type MovieInfo struct {
	Title  string
	Poster string
}

func downImage(url, fileName string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("无法下载图片", err)
		return
	}

	defer resp.Body.Close()

	// 创建本地文件用于保存图片
	file, err := os.Create(dir + fileName + ".jpg")
	if err != nil {
		fmt.Println("创建文件失败", err)
		return
	}

	defer file.Close()

	// 将图片数据写入本地文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("无法保存图片", err)
		return
	}

	fmt.Println("图片已成功下载和保存到本地文件", fileName)
}

func DownLoadPoster(title []string) {
	t := url.QueryEscape(strings.Join(title, " "))

	resp, err := http.Get(baseUrl + "&t=" + t)
	if err != nil {
		fmt.Println("查询电影信息失败", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("网络波动，请稍后再试", err)
		return
	}

	var result MovieInfo
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		fmt.Println("解析电影信息失败，请稍后再试", err)
		return
	}

	downImage(result.Poster, result.Title)
}
