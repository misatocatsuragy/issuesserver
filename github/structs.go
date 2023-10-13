// github page provides interface to GitHub API of quering to repo info:
// - users
// - ussues
// - milestones
package github

import "time"

const APIURL = "https://api.github.com"

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login         string
	HTMLURL       string `json:"html_url"`
	Contributions int
}

type Milestone struct {
	Number       int
	HTMLURL      string `json:"html_url"`
	Title        string
	State        string
	Description  string
	Creator      *User
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
}

type SearchResult struct {
	Issues     []*Issue
	Users      []*User
	Milestones []*Milestone
	RepoName   string
	OwnerName  string
	Error      string
}
