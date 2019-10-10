package main

import "fmt"

//!+编写一个rotate函数，通过一次循环完成旋转。
func main() {
	a := []int{0, 1, 2, 3, 4, 5}
	a = rotate2(a, 3)
	fmt.Println(a)
}

//!-

//!+ 要rotate的slice和rotate的位置
func rotate(a []int, rotateIndex int) []int {
	lens := len(a)
	if rotateIndex <= 0 && rotateIndex >= lens-1 {
		return a
	}
	arr := make([]int, lens)
	// 循环
	for i := range a {
		// 计算新数组第i个元素在原数据的索引
		index := rotateIndex + i
		if index >= lens {
			index %= lens
		}
		// 新数组
		arr[i] = a[index]
	}
	return arr
}

//利用切片的能力与内置copy方法
//将原数组rotateIndex开始到末尾的所有元素copy到新数组的开始到对应长度a
//将原数组开始到rotate index的所有元素copy到新数组的对应长度到结尾
func rotate2(a []int, rotateIndex int) []int {
	lens := len(a)
	if rotateIndex <= 0 && rotateIndex >= lens-1 {
		return a
	}
	arr := make([]int, lens)
	copy(arr[:lens-rotateIndex], a[rotateIndex:])
	copy(arr[lens-rotateIndex:], a[:rotateIndex])
	return arr
}
