package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func echo1() string {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "  "
	}

	return s
}

func echo2() {
	// var s,sep string
	for i, v := range os.Args {
		fmt.Println(i, v)
		fmt.Println()
	}
}

func dup1() {
	counts := make(map[string]int)
	// bufio.NewScanner()：读取输入并将其拆成行或单词
	input := bufio.NewScanner(os.Stdin)

	// Scan()：读入下一行
	for input.Scan() {
		// 输入""e"代表结束
		// 或者快捷键control + d
		// if input.Text() == "e" {
		// 	break
		// }
		// Text()：得到读取内容
		counts[input.Text()]++
		fmt.Println(counts)
	}

	for line, n := range counts {
		if n > 1 {
			// Printf()的第一个参数是格式字符串，指定后续参数被如何格式化
			// %d          十进制整数
			// %x, %o, %b  十六进制，八进制，二进制整数。
			// %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
			// %t          布尔：true或false
			// %c          字符（rune） (Unicode码点)
			// %s          字符串
			// %q          带双引号的字符串"abc"或带单引号的字符'c'
			// %v          变量的自然形式（natural format）
			// %T          变量的类型
			// %%          字面上的百分号标志（无操作数）
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	fmt.Println(files)

	if len(files) == 0 {
		// map是引用传递，不是值传递
		countLines(os.Stdin, counts)
	} else {
		for _, file := range files {
			f, err := os.Open(file)

			fmt.Println(f)

			//：nil === null
			if err != nil {
				// Fprintf：错误输出
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// ReadFile()返回字节切片（byte slice）
		// bufio.NewScanner()、ioutil.ReadFile()、ioutil.WriteFile()都使用*os.File
		//的Read和Write方法
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			fmt.Println(line)
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if counts[input.Text()] > 0 {
			fmt.Println(f.Name())
		}

		counts[input.Text()]++
	}
}

func testMap(anyMap map[string]int) {
	anyMap["q"] = 999
}

func testGet() {
	url := "https://xxx.com"
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	// 1.2
	// s:= echo1()
	// fmt.Println(s)
	// echo2()

	// dup1()
	// dup2()
	// dup3()

	// testCounts := make(map[string]int)
	// testMap(testCounts)
	// fmt.Println(testCounts)

	// rand.Seed(time.Now().UTC().UnixNano())
	// lissajous(os.Stdout)
	// create, _ := os.Create("out.gif")
	// lissajous(create)

	fetch()
	// fetchAll()
	// testGoroutine()
	// server1()
	// server2()
	// server3()
	// lissajousServer()
	// testGet()
}

func lissajous(out io.Writer, cycles int) {
	// var palette = []color.Color{color.White, color.Black}
	var palette = []color.Color{color.Opaque,
		color.RGBA{
			R: 0x00,
			G: 0xff,
			B: 0x00,
			A: 0xff,
		}}

	// 常量声明
	//：包级别共享
	//：必须是数字/字符串/固定的boolean
	const (
		whiteIndex = 0
		blackIndex = 1
	)

	const (
		// cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	// gif.GIF类型的struct
	// 初始值以外，类型包含的其他值为zero
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(blackIndex))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}

func fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = "https://" + url
		}

		fmt.Println(url)

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(resp.StatusCode)
		}

		b, err := ioutil.ReadAll(resp.Body)
		// _, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)

	}
}

func fetchAll() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetchN(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	// sets := <-ch
	// fmt.Println(sets)/

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetchN(url string, ch chan<- string) {
	start := time.Now()

	// resp, err := http.Get(url)
	resp, _ := http.Get(url)
	// if err != nil {
	// 	ch <- fmt.Sprint(err) // send to channel ch
	// 	return
	// }

	// 创建文件
	utl_file_name := url + strconv.Itoa(1) + ".txt"
	file, _ := os.Create(utl_file_name)

	// nbytes, _ := io.Copy(ioutil.Discard, resp.Body)
	nbytes, _ := io.Copy(file, resp.Body)
	resp.Body.Close() // don't leak resource
	// if err != nil {
	// 	ch <- fmt.Sprintf("while reading %s: %v", url, err)
	// 	return
	// }

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
	// stop

	fmt.Println("rest operations")
}

func testGoroutine() {
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetchN(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println("just print")
		fmt.Println(<-ch) // receive from channel ch
	}

	fmt.Println("end")
}

func server1() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func server2() {
	var mu sync.Mutex
	var count int

	// handler echoes the Path component of the requested URL
	handler := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	}

	// counter echoes the number of calls so far
	counter := func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		fmt.Fprintf(w, "Count: %d\n", count)
		mu.Unlock()
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8008", nil))
}

func server3() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}

		fmt.Fprintf(w, "Host = %q\n", r.Host)
		fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		for k, v := range r.Form {
			fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
		}
	}

	http.HandleFunc("/server3", handler)
	log.Fatal(http.ListenAndServe("localhost:8009", nil))
}

func lissajousServer() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Query", r.URL.Query())
		fmt.Println("Query cycles", r.URL.Query()["cycles"])
		cycles, err := strconv.Atoi(r.URL.Query()["cycles"][0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "cycles: %v\n", err)
		}

		lissajous(w, cycles)
	}

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func brief() {
	// 命名类型 === （ts）类型别名

	// 指针
	//：&，对变量操作，返回内存地址，类型为*int
	//：*，对指针操作，返回值

	// 方法和接口

	// 包（package）
}
