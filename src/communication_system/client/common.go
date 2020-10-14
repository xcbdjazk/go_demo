package main

import (
	"fmt"
	"net"

	"goproject/src/communication_system/model"
	"goproject/src/communication_system/utils"
)

func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:8081")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer conn.Close()
	var user model.User
	fmt.Print("name")
	_, _ = fmt.Scan(&user.UserName)
	fmt.Print("password")
	_, _ = fmt.Scan(&user.PassWord)
	datas, err := user.Marshal()
	if err != nil {
		fmt.Println("user.Marshal_err", err)
		return
	}
	err = utils.WriteMsg(conn, datas)
	if err != nil {
		fmt.Println("WriteMsg", err)
		return
	}
	msgStr, err := utils.ReadMsg(conn)
	fmt.Println(msgStr)
}
