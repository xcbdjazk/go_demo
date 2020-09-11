package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	n := time.Now()
	n = n.UTC()

	s := n.Format("2006-01-02 15:04:05 Z07:00")
	fmt.Println(s)
	fmt.Println(n)
	// 时间戳 毫秒
	ss := n.Unix()
	// 时间戳 纳秒 int64
	ss = n.UnixNano()
	fmt.Printf("%s, %T, %v \n", ss, ss, ss)

	// 时间戳转time字符串
	var sStr = strconv.FormatInt(ss, 10)
	fmt.Println("s_str>>", sStr)
	ss, _ = strconv.ParseInt(sStr, 10, 64)
	n = time.Unix(0, ss)
	s = n.Format("2006-01-02 15:04:05")
	fmt.Println(s)
	// time字符串 直接转为时间对象
	n, _ = time.ParseInLocation("2006-01-02 15:04:05", s, time.Local)
	fmt.Println(n)
	fmt.Println(time.Hour.String())

	//n.Add()

}
