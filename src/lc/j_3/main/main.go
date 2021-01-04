package main

import "fmt"

func replaceSpace(s string) string {
	var a []rune
	for _, i := range s {
		if string(i) == " " {
			a = append(a, '%', '2', '0')
		} else {
			a = append(a, i)
		}
	}
	return string(a)
}

func main() {
	fmt.Println(replaceSpace("1 2 3"))
}
