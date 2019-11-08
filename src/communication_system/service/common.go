package main

import (
	"communication_system/utils"
	"fmt"
	"net"
)

func doSomeThing(con net.Conn) {
	defer con.Close()
	str, err := utils.ReadMsg(con)
	if err != nil {
		fmt.Println("ReadMsg error ", err)
		return
	}
	fmt.Println(str)
	name := "hihihihihi"
	err = utils.WriteMsg(con, name)
	if err != nil {
		fmt.Println("ReadMsg error ", err)
		return
	}

}
func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("err__Listen", err)
		return
	}
	defer listener.Close()
	for {
		con, err := listener.Accept()
		if err != nil {
			fmt.Println("err__Accept:", err)
			return
		}
		go doSomeThing(con)
	}
}
