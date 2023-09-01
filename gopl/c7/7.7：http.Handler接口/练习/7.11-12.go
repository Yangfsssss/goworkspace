package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

// 练习 7.11： 增加额外的handler让客户端可以创建，读取，更新和删除数据库记录。
// 例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。
// （注意：这个修改会引入变量同时更新的问题）

type dollars float32

type database struct {
	mu    sync.Mutex
	items map[string]dollars
}

// func (db *database) list(w http.ResponseWriter, req *http.Request) {
// 	db.mu.Lock()
// 	defer db.mu.Unlock()

// 	for item, price := range db.items {
// 		fmt.Fprintf(w, "%s: %f\n", item, price)
// 	}
// }

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")

	db.mu.Lock()
	defer db.mu.Unlock()

	price, ok := db.items[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}

	fmt.Fprintf(w, "%f\n", price)
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item already exists: %q\n", item)
		return
	}

	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	db.items[item] = dollars(f)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "item created: %s:%f", item, db.items[item])
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.items[item]; !ok {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "item does not exist: %q\n", item)
		return
	}

	f, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "invalid price: %q\n", price)
		return
	}

	db.items[item] = dollars(f)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "price updated: %q %f\n", item, db.items[item])
}

func main() {
	launchServer()
}

func launchServer() {
	db := database{items: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// 练习 7.12： 修改/list的handler让它把输出打印成一个HTML的表格而不是文本。html/template包（§4.6）可能会对你有帮助。

const HTMLTemplate = `
<table>
	<tr>
		<th>Item</th>
		<th>Price</th>
	</tr>
{{range $item, $price := .}}
<tr>
	<td>{{$item}}</td>
	<td>{{$price}}</td>
</tr>
	{{end}}
</table>
`

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	db.mu.Lock()
	defer db.mu.Unlock()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl, err := template.New("table").Parse(HTMLTemplate)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, db.items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
