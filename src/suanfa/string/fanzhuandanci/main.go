package main

import "fmt"

// 查找子串 "google" in "good google"

func f(str2 string) string {
	var str []rune
	index := 0
	for i, v := range str2 {

		if string(v) == " " {
			fmt.Println(string([]rune(str2)[index:i]))
			str = append([]rune{' '}, str...)
			str = append([]rune(str2)[index:i], str...)
			index = i + 1
		}
	}
	str = append([]rune{' '}, str...)
	str = append([]rune(str2)[index:], str...)
	return string(str)
}

func init() {
	fmt.Println(f(" as1d ddd as2d 1a1 "))
}

func main() {

}
