package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func f1(d int, a ...interface{}) {
	fmt.Println(d)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	fmt.Printf("%#v, %T ", a, a)
}

type a1 struct {
	A func(int, ...interface{})
}

func RandSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	if rv.Type().Kind() != reflect.Slice {
		return
	}

	length := rv.Len()
	if length < 2 {
		return
	}

	swap := reflect.Swapper(slice)
	rand.Seed(time.Now().Unix())
	for i := length / 2; i >= 0; i-- {
		j := rand.Intn(length)
		swap(i, j)
	}
	return
}

func main() {
	//var a = []interface{}{
	//	"dd", 1, 'd',
	//}
	//f1(1, a[:3]...)
	//var a a1
	//a.A = f1
	//a.A(1)
	var a = []int{1, 2, 3, 4, 5}

	RandSlice(a)
	fmt.Println(a)
}
