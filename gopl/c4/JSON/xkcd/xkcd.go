package xkcd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
)

const baseUrl = "https://xkcd.com"

const urlSuffix = "/info.0.json"

const dir = "/Users/bytedance/Documents/code/goworkspace/gopl/c4/JSON/xkcd/csv/"

type Comic struct {
	Title string `json:"safe_title"`
	Img   string `json:"img"`
	Link  string `json:"link"`
}

func GenerateOfflineComicIndex() {
	data := make(map[int]Comic)

	for i := 1; i < 30; i++ {
		c, err := http.Get(baseUrl + "/" + strconv.Itoa(i) + urlSuffix)
		if err != nil {
			continue
		}

		defer c.Body.Close()

		if c.StatusCode != http.StatusOK {
			break
		}

		var result Comic
		if err := json.NewDecoder(c.Body).Decode(&result); err != nil {
			fmt.Println("解析失败，请稍后再试", err)
			continue
		}

		data[i] = result
	}

	file, err := os.Create(dir + "comic.csv")
	if err != nil {
		fmt.Println("创建csv文件失败", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	// 获取map的键，并对键进行排序
	keys := make([]int, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// 写入表头
	headers := []string{"#", "title", "img", "link"}
	writer.Write(headers)

	// 按排序后的键顺序遍历map，写入键值对
	for _, key := range keys {
		v := data[key]
		row := []string{strconv.Itoa(key), v.Title, v.Img, v.Link}
		writer.Write(row)
	}

	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("写入csv文件失败", err)
		return
	}

	fmt.Println("数据已成功写入到 comic.csv 文件")
}

func SearchComic(t []string) {
	input := strings.ToLower(strings.Join(t, " "))
	file, err := os.Open(dir + "comic.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)

	// 读取所有记录
	records, err := reader.ReadAll()
	if err == io.EOF {
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	// 检索数据
	for _, record := range records {
		title := record[1]
		img := record[2]

		if strings.Contains(strings.ToLower(title), input) {
			fmt.Printf("Title:%s\tImage:%s\t\n", title, img)
			return
		}
	}
}
