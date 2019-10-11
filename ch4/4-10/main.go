package main

import (
	"fmt"
	"log"
	"os"
	"test/gopl_exercises/github"
)

//!+修改issues程序，根据问题的时间进行分类，比如不到一个月的、不到一年的、超过一年。
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
