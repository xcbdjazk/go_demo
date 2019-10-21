package message

import "encoding/json"

type Message struct {
	Code    int    `json:"code"`
	Error   string `json:"error"`
	Message string `json:"message"`
}

func Marshal(s Message) (st string, err error) {
	ns, err := json.Marshal(s)
	st = string(ns)
	return
}

func UnMarshal(s string) (message Message, err error) {
	err = json.Unmarshal([]byte(s), &message)
	return
}
