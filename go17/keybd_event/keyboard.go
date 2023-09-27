package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	// 打开键盘事件监听。
	err := keyboard.Open()
	if err != nil {
		fmt.Println("Error opening keyboard:", err)
		return
	}
	defer keyboard.Close()

	fmt.Println("Listening for key presses. Press 'ESC' to exit.")

	// 循环监听键盘事件。
	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Println("Error getting key:", err)
			return
		}

		// 检测是否按下 'ESC' 键，如果是则退出循环。1123123
		if key == keyboard.KeyEsc {
			fmt.Println("Exiting...")
			break
		}

		// 输出按下的键信息。
		fmt.Printf("Key: %c, Code: %v\n", char, key)
	}
}
