package main

import "fmt"

// 查找子串 "google" in "good google"

func findStr(str1, str2 string) bool {
	if len(str1) > len(str2) {
		return false
	}
	for i := 0; i <= len([]rune(str2))-len([]rune(str1))+1; i++ {
		if []rune(str2)[i] == []rune(str1)[0] {
			jc := 0
			for j := 0; j < len([]rune(str1)); j++ {

				if []rune(str2)[i+j] != []rune(str1)[j] {
					break
				}
				jc = j
			}
			if jc == len([]rune(str1))-1 {
				return true
			}
		}

	}
	return false
}

func init() {
	fmt.Println(findStr("你是我的1", "你是我的"))
}

func main() {

}
