package github

//发起http请求,解析返回结果

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// SearchIssues 从github上搜索issue
func SearchIssues(terms []string) (*IssueSearchResult, error) {
	// 构造query param
	queryStr := strings.Join(terms, "+")
	queryParam := url.QueryEscape(queryStr)
	requestStr := SearchURL + "/issues?q=" + queryParam
	resp, err := http.Get(requestStr)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d\nsearch query failed case:%s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	var result IssueSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
