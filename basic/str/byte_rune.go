package main

import "fmt"

func init() {

}
func main() {
	// '' 表示字符型 字符型就是int类型
	//var a = "a"
	//fmt.Printf("s = %s, v =  %v", a, a)
	// " "表示字符串型 字符型就是string类型
	//var a = "a123"
	//fmt.Printf("s = %s, v =  %v, %T", a[2], a[2], a[2])
	// 汉字使用的是utf-8编码
	//var a = '你'
	//fmt.Printf("s = %s, v =  %v, %T %c", a, a, a, a)
	// 英文替换字符
	//var a = "ABC123"
	//newStr := []byte(a)
	//newStr[1] = '1'
	//fmt.Println(string(newStr))
	// 非英文替换字符
	//var a = "ABC123"
	//newStr := []rune(a)
	//newStr[1] = '我'
	//fmt.Println(string(newStr), len(newStr))

	// 遍历字符串
	var s rune
	a := "我十岁abc231"
	for _, s = range a {
		fmt.Println(string(s))
	}
	fmt.Println(s)

}
