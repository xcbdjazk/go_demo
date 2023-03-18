package main

import "fmt"

// 括号是否有效
var s = make([]string, 4, 8)

func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func pop(str string) bool {

	for _, v := range str {
		if IsContain([]string{"{", "(", "["}, string(v)) {
			s = append(s, string(v))
		} else {
			p := s[len(s)-1]
			if string(v) == "}" && p != "{" {
				return false
			}
			if string(v) == "]" && p != "[" {
				return false
			}
			if string(v) == ")" && p != "(" {
				return false
			}
			s = s[:len(s)-1]
		}
	}
	return len(s) == 0
}

func init() {
	var str = "{}{}{[}{]}{}"
	fmt.Println(pop(str))
}

func main() {

}
