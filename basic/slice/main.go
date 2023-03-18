package main

import "fmt"

func main() {
	make_qiepian()
}

func shuzu() {
	// 申明数组的方法
	// 1.
	//var a [3]int
	// a[0] = 33
	//fmt.Println(a)
	// 2.
	//var a = [3]int{1,2,3}
	//fmt.Println(a)
	// 3.
	//var a = [...]int{1,2,3}
	//fmt.Println(a)

	// 4.
	var a = [...]int{1: 1, 2: 2, 3: 3}
	fmt.Printf("%v, %T, 长度%v", a, a, len(a))

}

func qiepian() {
	//var a []int
	//a == nil // true
	//fmt.Printf("%v, %T, 长度%v", a, a,len(a))
	//var a = []int{1,2,3,4}
	//fmt.Printf("%v, %T, 长度%v", a, a,len(a))
	// 基于数组创建切片
	var a = [4]int{1, 2, 3, 4}
	b := a[:3]
	fmt.Printf("%v, %T, 长度%v \n", b, b, len(b))
	c := b[1:3]
	b[1] = 55
	// 查看容量
	// cap()
	// 切片的容量是其第一个元素开始,到其底层数组的最后一个元素末尾的个数
	// b.cap == len a[1:]
	fmt.Printf("%v, %T, 长度%v, 容量%v", c, c, len(c), cap(c))

}

func make_qiepian() {
	var a = make([]int, 4, 8)
	a[1] = 33
	fmt.Printf("%v, %T, 长度%v, 容量%v \n ", a, a, len(a), cap(a))

	a = append(a, 1, 2, 3, 4, 5, 6)

	fmt.Printf("%v, %T, 长度%v, 容量%v  \n ", a, a, len(a), cap(a))
	// 复制切片
	var b = make([]int, len(a), cap(a))
	copy(b, a)
	fmt.Printf("%v, %T, 长度%v, 容量%v  \n ", b, b, len(b), cap(b))

}
