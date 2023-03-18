package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	abac()
}

type P struct {
	Age  int `json:"age"`
	Name string
}

func abac() {
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		panic(err)
	}

	// 添加自定义函数
	e.AddFunction("check_name", func(arguments ...interface{}) (interface{}, error) {
		if len(arguments) == 0 {
			return false, nil
		}
		p := arguments[0].(*P)

		if p.Name == "123" {
			return true, nil
		}
		return false, nil
	})

	fmt.Println(e.GetPolicy())
	{

		ok, err := e.Enforce(&P{
			Age:  19,
			Name: "123",
		}, "/data1", "read")
		if err != nil {
			panic(err)
		}
		fmt.Println(ok)
	}

	// 添加一个规则
	ok, err := e.AddPolicy("false", "/data2", "write")
	if err != nil {
		panic(err)
	}
	fmt.Println("AddPolicy, ok", ok)
	fmt.Println(e.GetPolicy())
	{

		ok, err := e.Enforce(&P{
			Age:  129,
			Name: "123",
		}, "/data2", "write")
		if err != nil {
			panic(err)
		}
		fmt.Println(ok)
	}
	// 保存规则
	err = e.SavePolicy()
	if err != nil {
		panic(err)
	}

}
