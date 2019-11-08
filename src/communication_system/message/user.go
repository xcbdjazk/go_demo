package message

import "encoding/json"

type User struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

func (s *User) Marshal() (st string, err error) {
	ns, err := json.Marshal(s)
	st = string(ns)
	return
}
