package github

import (
	"html/template"
	"io"
)

const mainPageTemplate = `
<h1>{{.OwnerName}}/{{.RepoName}}</h1>
<ul>
<li><a href='http://localhost:8080/issues'>issues</li>
  <li><a href='http://localhost:8080/users'>users</li>
  <li><a href='http://localhost:8080/milestones'>milestones</li>
</ul>`

const errorPageTemplate = `
<h1>ERROR</h1>
<p>Request failed:{{.Error}}</p>
`

const issuesPageTemplate = `
<h1>{{len .Issues}} issues</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>User</th>
  <th>Title</th>
</tr>
{{range .Issues}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`

const usersPageTemplate = `
<h1>{{len .Users}} users</h1>
<table>
<tr style='text-align: left'>
  <th>Login</th>
  <th>Contributions</th>
</tr>
{{range .Users}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
  <td>{{.Contributions}}</td>
</tr>
{{end}}
</table>
`

const milestonesPageTemplate = `
<h1>{{len .Milestones}} milestones</h1>
<table>
<tr style='text-align: left'>
  <th>#</th>
  <th>State</th>
  <th>Creator</th>
  <th>Title</th>
  <th>Open issues</th>
  <th>Close issues</th>
</tr>
{{range .Milestones}}
<tr>
  <td><a href='{{.HTMLURL}}'>{{.Number}}</td>
  <td>{{.State}}</td>
  <td><a href='{{.Creator.HTMLURL}}'>{{.Creator.Login}}</a></td>
  <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
  <td>{{.OpenIssues}}</td>
  <td>{{.ClosedIssues}}</td>
</tr>
{{end}}
</table>
`

func GenPageHTML(out io.Writer, resource string, sr *SearchResult) error {
	var page *template.Template

	switch resource {
	case "main":
		page = template.Must(template.New("main").Parse(mainPageTemplate))
	case "error":
		page = template.Must(template.New("error").Parse(errorPageTemplate))
	case "issues":
		page = template.Must(template.New("issues").Parse(issuesPageTemplate))
	case "users":
		page = template.Must(template.New("users").Parse(usersPageTemplate))
	case "milestones":
		page = template.Must(template.New("milestones").Parse(milestonesPageTemplate))
	}

	if err := page.Execute(out, sr); err != nil {
		return err
	}

	return nil
}
