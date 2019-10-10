package main

import "fmt"

//!+重写reverse函数，使用数组指针代替slice。
func main() {
	a := [...]int{0, 1, 2, 3}
	// gi方法参数传递是值传递,想要通过方法修改原数组，要么传递指针，要么传递slice。
	// 传递指针时方法参数签名的数组声明的长度必须与实参长度一致
	p := &a
	reverse(p)
	fmt.Println(a)
}

func reverse(s *[4]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-
