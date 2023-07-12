package main

import (
	"github"
	"github_info"
	"html/template"
	"log"
	"os"
	"time"
)

const templ = `
{{.TotalCount}} issues:
{{range .Items}}---------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreateAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

var report = template.Must(template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func textAndHTMLTemplate() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// if err := report.Execute(os.Stdout, result); err != nil {
	// 	log.Fatal(err)
	// }
	if err := issueList.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

var issueList = template.Must(template.New("issueList").Parse(`
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

func safeHTML() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	t := template.Must(template.New("safeHTML").Parse(templ))

	var data struct {
		A string
		B template.HTML
	}

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"

	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

func issueHtml() {}

func main() {
	// textAndHTMLTemplate()
	// safeHTML()
	github_info.GetGithubInfo()
}
