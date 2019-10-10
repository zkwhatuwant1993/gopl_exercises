package main

//!+写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func main() {
	s := []string{"aa", "bb", "bb", "c"}
	s = removeAdjacentDuplicated(s)
}

//!-

// 原地消除必须在原数组上完成操作（即不能新分配内存）
// 遇到相同的元素，所有其后的元素往前覆盖
func removeAdjacentDuplicated(s []string) []string {
	var (
		lens        = len(s)
		removeCount = 0
		cap         = 0
	)
	for i := range s {
		if s[i] == s[i+i] {
			copy(s[i:], s[i+1:])
			s = s[i:len(s)-1]
		}
	}
}
