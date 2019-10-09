package main

import (
	"fmt"
	"strconv"
	"strings"
)

//!+练习3-10与3-11 编写一个非递归版本的comma函数，使用bytes.Buffer代替字符串链接操作。

var (
	a = 1.2345
	b = 1234.5678
	c = -1234.5678
)

func main() {
	fmt.Println(commaR(num2string(a)))
	fmt.Println(commaR(num2string(b)))
	fmt.Println(commaR(num2string(c)))
}

func num2string(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func commaR(s string) string {
	// 是否是负数
	isNeg := strings.Contains(s, "-")
	if isNeg {
		s = s[1:]
	}
	// 获取小数点位置
	var result string
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		result = commaIntPart(s[:dot]) + "." + commaDecPart(s[dot+1:])
	} else {
		result = commaIntPart(s)
	}
	if isNeg {
		result = "-" + result
	}
	return result
}

// Int Part
func commaIntPart(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return commaIntPart(s[:n-3]) + "," + s[n-3:]
}

// Decimal Part
func commaDecPart(s string) string {
	if len(s) <= 3 {
		return s
	}
	return s[:3] + "," + commaDecPart(s[3:])
}
