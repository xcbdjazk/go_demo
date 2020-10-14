package main

import (
	"fmt"
	"goproject/src/practice1/account"
	"goproject/src/practice1/user"
)

func main() {
	var point string
	fmt.Println("-------银行存取款系统V1.0------")
	flag := true
	for flag {
		fmt.Println("-------功能区，1.登录系统，2.注册，3，退出------")
		fmt.Print("请输入:")
		fmt.Scan(&point)
		fmt.Println("")
		switch point {
		case "1":
			nowUser, ok := login()
			if ok {
				loginDo(nowUser)
			}
		case "2":
			register()
		case "3":
			flag = false
			fmt.Println("你已退出")
		case "99":
			fmt.Println(user.Users)
		default:
			fmt.Println("输入无效")
		}
	}

}
func login() (user.User, bool) {
	var name string
	var password string
	count := 0
lable1:
	if count == 3 {
		fmt.Println("超出登录次数")
		return user.User{}, false
	}
	fmt.Print("请输入用户名称:")
	fmt.Scan(&name)
	fmt.Println("")
	fmt.Print("请输入用户密码:")
	fmt.Scan(&password)
	nowUser, ok := user.GetUser(name, password)
	if ok {
		fmt.Println("登录成功")
		return nowUser, true
	} else {
		count++
		fmt.Println("登录失败")
		goto lable1
	}
}
func loginDo(nowUser user.User) {
lable:
	flag := true
	for flag {
		fmt.Println("---------0.查询余额，1.取钱，2.存钱，3.修改密码，4.返回上一层")
		var point string
		fmt.Print("请输入:")
		fmt.Scan(&point)
		switch point {
		case "0":
			au := account.GetAU(&nowUser)
			fmt.Println("你的余额：", au.Money)
			break
		case "1":
			money := 0.0
			fmt.Print("请输入金额:")
			fmt.Scan(&money)
			account.GetMoney(&nowUser, money)
		case "2":
			money := 0.0
			fmt.Print("请输入金额:")
			fmt.Scan(&money)
			account.SaveMoney(&nowUser, money)
		case "3":
			var newPassWord string
			fmt.Print("请输入新密码:")
			fmt.Scan(&newPassWord)
			nowUser.SetUserPassWord(newPassWord)
			fmt.Println("修改成功")
			goto lable
		case "4":
			flag = false
		default:
			fmt.Println("输入无效")
		}
	}

}
func register() {
	var name string
	var password string
	conut := 0
lable2:
	if conut == 3 {
		fmt.Println("超过操作次数")
		return
	}
	fmt.Print("请输入用户名称:")
	fmt.Scan(&name)
	fmt.Println("")
	fmt.Print("请输入用户密码:")
	fmt.Scan(&password)
	_, ok := user.AddUser(name, password)
	if ok {
		fmt.Println("注册成功")
	} else {
		conut++
		fmt.Println("注册失败，用户名已经存在")
		goto lable2
	}
}
