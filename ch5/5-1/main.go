package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

//!+ 修改findlinks代码中遍历n.FirstChild链表的部分，将循环调用visit，改成递归调用。
func main() {
	node, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, v := range visit(nil, node) {
		fmt.Println(v)
	}
}

//!-

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, a := range node.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 分析原循环的执行：深度优先搜索
	// 1.解析当前节点本身
	// 2. 是否有子节点，开始递归解析 回到1
	// 3. 如果有兄弟节点，开始递归解析 回到1
	// 即沿着doc树的某一分支，递归到底，然后再从分支最远端的节点的兄弟节点重复这一过程
	// child := node.FirstChild; child != nil; child = child.NextSibling {
	//	links = visit(links, child)
	//}

	if node.FirstChild != nil {
		links = visit(links, node.FirstChild)
	}
	if node.NextSibling != nil {
		links = visit(links, node.NextSibling)
	}

	return links
}
