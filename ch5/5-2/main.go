package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
)

var result map[string]int = make(map[string]int)

// 编写函数，记录在HTML树中出现的同名元素的次数
func main() {
	node, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("5-1 parse error:%v\n", err)
	}
	countElement(node)
	fmt.Printf("%-10s %-10s\n", "tag", "count")
	for k, v := range result {
		fmt.Printf("%-10s %-10d\n", k, v)
	}
}

func countElement(node *html.Node) {
	if node.Type == html.ElementNode {
		result[node.Data]++
	}
	if node.FirstChild != nil {
		countElement(node.FirstChild)
	}
	if node.NextSibling != nil {
		countElement(node.NextSibling)
	}
}
