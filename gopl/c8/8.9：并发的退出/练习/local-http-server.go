package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Data struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Age   int     `json:"age"`
	City  string  `json:"city"`
	Money float64 `json:"money"`
}

func main() {
	http.HandleFunc("/", getData)
	http.ListenAndServe(":8899", nil)
	fmt.Println("http server running")
}

func getData(w http.ResponseWriter, r *http.Request) {
	// 生成随机数据
	rand.Seed(time.Now().UnixNano())
	data := Data{
		ID:    rand.Intn(1000),
		Name:  "John Doe",
		Age:   rand.Intn(50) + 20,
		City:  "New York",
		Money: float64(rand.Intn(1000000)) / 100,
	}

	// 将数据转换为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Simple API")

	idStr := r.URL.Path[len("/data/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "无效的数据 ID", http.StatusBadRequest)
		return
	}

	fmt.Println("收到请求，id为：", id)

	// 发送响应
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
