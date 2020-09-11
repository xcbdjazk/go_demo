package utils

import (
	"encoding/binary"
	"fmt"
	"net"
)

func ReadMsg(con net.Conn) (str string, err error) {
	var b1 []byte = make([]byte, 8096)

	_, err = con.Read(b1[:4])
	if err != nil {
		fmt.Println("con__Read,", err)
		return
	}
	var pagLength uint32
	pagLength = binary.BigEndian.Uint32(b1[0:4])

	n, err := con.Read(b1[:pagLength])
	if uint32(n) != pagLength || err != nil {
		fmt.Println("errr", err)
		return
	}

	str = string(b1[:pagLength])
	return
}
