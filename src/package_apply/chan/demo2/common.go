package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {
	var interfaceChan chan interface{}
	interfaceChan = make(chan interface{}, 3) // 创建管道的长度，不能超过 否则会死锁
	//向管道写入数据
	interfaceChan <- Cat{Name: "mao1", Age: 12}
	interfaceChan <- Cat{Name: "mao2", Age: 13}
	interfaceChan <- -1
	//管道取出数据
	//inter:= <- interfaceChan
	//fmt.Println(inter.Name) // error inter为interface{} 类型，应该使用类型推导
	//
	inter, _ := (<-interfaceChan).(Cat) // len <= 0  会死锁
	fmt.Println(inter.Name)
	fmt.Printf("长度（len）%v, 容量（cap）%v \n ", len(interfaceChan), cap(interfaceChan))
	close(interfaceChan)
	//close(interfaceChan) //  不能重复关闭
	//interfaceChan <- -1  // 关闭后不能再写入
	for k := range interfaceChan { //关闭后才能使用for range 取出
		fmt.Printf("%v \n ", k)

	}
}
