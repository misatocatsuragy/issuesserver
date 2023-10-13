package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Query to owner/repo
func GetRepoInfo(owner string, repo string) (*SearchResult, error) {
	result := &SearchResult{OwnerName: owner, RepoName: repo}
	issues, err := getIssues(owner, repo)
	if err != nil {
		return result, err
	}
	users, err := getUsers(owner, repo)
	if err != nil {
		return result, err
	}
	milestones, err := getMilestones(owner, repo)
	if err != nil {
		return result, err
	}

	result.Issues = issues
	result.Users = users
	result.Milestones = milestones

	return result, nil
}

// Get issues
func getIssues(owner string, repo string) ([]*Issue, error) {
	path := APIURL + "/repos/" + owner + "/" + repo + "/issues"
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error while quering repo info: %s", resp.Status)
	}

	var result []*Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return result, nil
}

// Get users
func getUsers(owner string, repo string) ([]*User, error) {
	path := APIURL + "/repos/" + owner + "/" + repo + "/contributors"
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error while quering repo info: %s", resp.Status)
	}

	var result []*User
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return result, nil
}

// Get milestones
func getMilestones(owner string, repo string) ([]*Milestone, error) {
	path := APIURL + "/repos/" + owner + "/" + repo + "/milestones"
	resp, err := http.Get(path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("error while quering repo info: %s", resp.Status)
	}

	var result []*Milestone
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	resp.Body.Close()
	return result, nil
}
