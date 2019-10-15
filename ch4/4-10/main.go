package main

import (
	"fmt"
	"log"
	"os"
	"test/gopl_exercises/github"
	"time"
)

//!+修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// 计算时间
	now := time.Now()
	before1month := now.AddDate(0, -1, 0)
	before1year := now.AddDate(-1, 0, 0)

	// 使用map保存各时间段数据
	var itemMap map[int][]github.Issue
	itemMap = make(map[int][]github.Issue, 3)
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.Sub(before1month).Milliseconds() > 0 {
			itemMap[0] = append(itemMap[0], item)
		} else if item.CreatedAt.Sub(before1year).Milliseconds() > 0 {

			itemMap[1] = append(itemMap[1], item)
		} else {

			itemMap[2] = append(itemMap[2], item)
		}
	}
	dateStr := []string{"不到一月", "不足一年", "超过一年"}
	for k, v := range itemMap {
		fmt.Println(dateStr[k])
		for _, item := range v {
			fmt.Printf("#%-5d %9.9s %.55s %s %s \n",
				item.Number, item.User.Login, item.Title, item.CreatedAt, now.Sub(item.CreatedAt))
		}

	}

}
