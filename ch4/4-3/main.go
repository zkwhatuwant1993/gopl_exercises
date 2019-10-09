package main

import "fmt"

//!+重写reverse函数，使用数组指针代替slice。
func main() {
	a := []int{0, 1, 2, 3}
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[]int) {
	
}

//!-
