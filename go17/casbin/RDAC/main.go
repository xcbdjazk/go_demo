package main

import (
	"fmt"
	"github.com/casbin/casbin/v2"
)

func main() {
	abac()
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
		return false, nil
	})

	fmt.Println(e.GetPolicy())
	{

		ok, err := e.Enforce("abc", "", "/data1", "read")
		if err != nil {
			panic(err)
		}
		fmt.Println(ok)
	}

	// 添加一个规则
	//ok, err := e.AddPolicy("admin1", "domain3", "/data3", "read")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("AddPolicy, ok", ok)
	//fmt.Println(e.GetPolicy())
	//
	//fmt.Println("AddPolicy, ok", ok)
	//ok, err = e.AddRolesForUser("123", []string{"admin3"}, "domain3")
	//if err != nil {
	//	panic(err)
	//}

	//ok, err = e.AddNamedGroupingPolicies("g2", [][]string{
	//	{"123", "admin3", "domain3"},
	//})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("AddNamedGroupingPolicies, ok", ok)

	{

		ok, err := e.Enforce("123", "domain3", "/data3", "read")
		if err != nil {
			panic(err)
		}
		fmt.Println(ok)
	}
	// 保存规则
	//err = e.SavePolicy()
	//if err != nil {
	//	panic(err)
	//}

}
