package main

import (
	"fmt"
	"unicode"
)

//!+编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func main() {
	s := []byte("a bg  c	 d	e")
	s = removeAdjacentDuplicated(s)
	fmt.Println(string(s))
}

//!-

func removeAdjacentDuplicated(s []byte) []byte {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) &&
			unicode.IsSpace(rune(s[i+1])) {
			// 向前覆盖
			copy(s[i:], s[i+1:])
			// 将新的slice view更亲给切片
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}
