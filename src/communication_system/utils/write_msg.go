package utils

import (
	"encoding/binary"
	"fmt"
	"net"
)

func WriteMsg(con net.Conn, data string) (err error) {
	var pkgLength uint32
	pkgLength = uint32(len(data))
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], pkgLength)
	_, err = con.Write(b[:])
	if err != nil {
		fmt.Println("err___Write", err)
		return
	}
	_, err = con.Write([]byte(data))
	return err
}
