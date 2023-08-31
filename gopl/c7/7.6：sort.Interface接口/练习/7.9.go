package main

import (
	"log"
	"net/http"
	"sort"
	"text/template"
	"time"
)

// 练习 7.9： 使用html/template包（§4.6）替代printTracks将tracks展示成一个HTML表格。
// 将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	// time.ParseDuration用于将字符串表示的时间段解析为 time.Duration 类型的值
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return d
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

const HTMLTemplate = `
<table>
    <thead>
        <tr>
            <th><a href="/sort?column=Title">Title</a></th>
            <th><a href="/sort?column=Artist">Artist</a></th>
            <th><a href="/sort?column=Album">Album</a></th>
            <th><a href="/sort?column=Year">Year</a></th>
            <th><a href="/sort?column=Length">Length</a></th>
        </tr>
    </thead>
    <tbody>
        {{range .}}
        <tr>
            <td>{{.Title}}</td>
            <td>{{.Artist}}</td>
            <td>{{.Album}}</td>
            <td>{{.Year}}</td>
            <td>{{.Length}}</td>
        </tr>
        {{end}}
    </tbody>
</table>
`

func sortTracksByColumn(column string) {
	switch column {
	case "Title":
		// sort.Slice用于对切片进行排序。它的作用是根据指定的排序规则对切片进行排序，并且可以通过传递一个 less 函数来自定义排序规则
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Title < tracks[j].Title
		})
	case "Artist":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Artist < tracks[j].Artist
		})
	case "Album":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Album < tracks[j].Album
		})
	case "Year":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Year < tracks[j].Year
		})
	case "Length":
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Length < tracks[j].Length
		})
	}
}

func renderTracks(w http.ResponseWriter) {
	tmpl, err := template.New("table").Parse(HTMLTemplate)
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, tracks)
	if err != nil {
		log.Fatal(err)
	}
}

func sortHandler(w http.ResponseWriter, req *http.Request) {
	column := req.URL.Query().Get("column")
	sortTracksByColumn(column)
	renderTracks(w)
}

func main() {
	http.HandleFunc("/sort", sortHandler)
	http.ListenAndServe(":8080", nil)
}
