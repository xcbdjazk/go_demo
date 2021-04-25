package main

import (
"fmt"
	"unsafe"
)


func main1() {
	u := uint32(32)
	i := int32(1)
	fmt.Println(&u, &i) // 打印出地址
	p := &i             // p 的类型是 *int32
	//p = &u              // &u的类型是 *uint32，于 p 的类型不同，不能赋值
	//p = (*int32)(&u)    // 这种类型转换语法也是无效的
	fmt.Println(p)
}

func main() {
	u := uint32(32)
	i := int32(1)
	i2 := int32(1)
	fmt.Println(&u, &i, &i2)
	p := &i
	fmt.Println(unsafe.Pointer(&u))
	p = (*int32)(unsafe.Pointer(&u))
	fmt.Println(p)
	fmt.Println(*p)

}