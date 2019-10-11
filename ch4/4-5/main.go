package main

import "fmt"

//!+写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func main() {
	s := []string{"aa", "bb", "bb", "bb", "c"}
	s = removeAdjacentDuplicated(s)
	fmt.Println(s)
}

//!-

// 原地消除必须在原数组上完成操作（即不能新分配内存）
// 遇到相同的元素，所有其后的元素往前覆盖
func removeAdjacentDuplicated(s []string) []string {
	// 不使用range s来循环的原因是slice s的长度一直在变
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[0 : len(s)-1]
			// 向前覆盖之后，继续从当前位置开始判断
			i--
		}
	}
	return s
}
