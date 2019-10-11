package github

import "time"

// SearchURL 是github api v3的search接口
const SearchURL = "https://api.github.com/search"

// User is the github user who publish the issue
type User struct {
	Login   string `json:"login"`
	HTMLURL string `json:"html_url"`
}

// An Issue is a place to discuss ideas, enhancements, tasks, and bugs for a project.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"create_at"`
	Body      string    // in Markdown format
}

// IssueSearchResult 表示http请求结果
type IssueSearchResult struct {
	TotalCount int
	Items      []Issue
}
