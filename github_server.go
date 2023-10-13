// github_server is a simple server, provides naviagtion
// of the list of bug reports, milestones and users
package main

import (
	"fmt"
	"log"
	"net/http"

	github "githubserver/github"
)

var result = &github.SearchResult{}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	var repo, owner string
	if err := r.ParseForm(); err != nil {
		result.Error = fmt.Sprintf("%v", err)
		if genErr := github.GenPageHTML(w, "error", result); genErr != nil {
			log.Fatalf("%v", genErr)
		}
		return
	}
	for k, v := range r.Form {
		if k == "repo" {
			repo = v[0]
			continue
		}
		if k == "owner" {
			owner = v[0]
			continue
		}
	}
	if len(owner) == 0 || len(repo) == 0 {
		result.Error = "Specify repo or owner in request"
		github.GenPageHTML(w, "error", result)
		return
	}
	var err error
	result, err = github.GetRepoInfo(owner, repo)
	if err != nil {
		log.Fatalf("%v", err)
	}

	if genErr := github.GenPageHTML(w, "main", result); genErr != nil {
		log.Fatalf("%v", genErr)
	}
}

func issuesHandler(w http.ResponseWriter, r *http.Request) {
	if genErr := github.GenPageHTML(w, "issues", result); genErr != nil {
		log.Fatalf("%v", genErr)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if genErr := github.GenPageHTML(w, "users", result); genErr != nil {
		log.Fatalf("%v", genErr)
	}
}

func milestonesHandler(w http.ResponseWriter, r *http.Request) {
	if genErr := github.GenPageHTML(w, "milestones", result); genErr != nil {
		log.Fatalf("%v", genErr)
	}
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/issues", issuesHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/milestones", milestonesHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
