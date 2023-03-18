package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 其他基本类型转字符串
	var i int64 = 64
	var f float64 = 32.3
	var t bool = true
	var b byte = 'a'

	str1 := fmt.Sprintf("%d", i)
	fmt.Printf("类型 %T 值 %v \n", str1, str1)

	str1 = fmt.Sprintf("%f", f)
	fmt.Printf("类型 %T 值 %v \n", str1, str1)

	str1 = fmt.Sprintf("%t", t)
	fmt.Printf("类型 %T 值 %v \n", str1, str1)

	str1 = fmt.Sprintf("%c", b)
	fmt.Printf("类型 %T 值 %v \n", str1, str1)

	// 数字转为
	// 10 十进制
	str1 = strconv.FormatInt(int64(i), 10)
	fmt.Printf("类型 %T 值 %v \n", str1, str1)

	// 浮点转为str
	str1 = strconv.FormatFloat(f, 'f', 4, 64)
	fmt.Printf("类型 %T 值 %v \n", str1, str1)
	fmt.Printf("类型 %T 值 %v \n", 'a', 'a')

	str := "1"
	// string 转 int
	i_, e := strconv.ParseInt(str, 10, 64)
	fmt.Printf("类型 %T 值 %v \n", e, e)
	fmt.Printf("类型 %T 值 %v \n", i_, i_)
	var a = func() bool {
		fmt.Printf("123")
		return true
	}
	if 1 < 2 && a() {
		fmt.Printf("123444")
	}
}
