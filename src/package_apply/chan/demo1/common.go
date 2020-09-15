package main

import "fmt"

func main() {
	var intChan chan int
	intChan = make(chan int, 3) // 创建管道的长度，不能超过 否则会死锁
	//向管道写入数据
	intChan <- -1
	fmt.Printf("长度（len）%v, 容量（cap）%v \n ", len(intChan), cap(intChan))
	//管道取出数据
	num, ok := <-intChan // len <= 0  会死锁
	fmt.Println(num)
	fmt.Println(ok)
	num, ok = <-intChan // len <= 0  会死锁
	fmt.Println(num)
	fmt.Println(ok)
	fmt.Printf("长度（len）%v, 容量（cap）%v \n ", len(intChan), cap(intChan))
}
