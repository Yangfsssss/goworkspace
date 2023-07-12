package github_info

import (
	"fmt"
	"github"
	"html/template"
	"log"
	"os"
)

var infoList = template.Must(template.New("infoList").Parse(`
<h1>{{.TotalCount}} Github Info</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>Bug report</th>
	<th>Milestone</th>
	<th>Userinfo</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
</tr>
{{end}}
</table>
`))

func GetGithubInfo() {
	result, err := github.SearchIssues(os.Args[1:])

	if err != nil {
		log.Fatal(err)
	}

	htmlFile, err := os.Create("info.html")
	if err != nil {
		log.Fatal(err)
	}

	defer htmlFile.Close()

	if err := infoList.Execute(htmlFile, result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(htmlFile)
}
