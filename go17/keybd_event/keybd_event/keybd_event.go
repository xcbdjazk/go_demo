package main

import (
	"fmt"
	"time"

	"github.com/micmonay/keybd_event"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		fmt.Println("Error creating keybonding:", err)
		return
	}

	// 在Windows上，调用SetKeys方法模拟Windows的音量快捷键（音量增加：VolumeUp，音量减少：VolumeDown，静音：VolumeMute）。
	// 在macOS上，调用SetKeys方法模拟macOS的音量快捷键（音量增加：SOUND_UP，音量减少：SOUND_DOWN，静音：SOUND_MUTE）。
	// 在Linux上，也可以使用该方法，但是快捷键可能需要根据Linux发行版的不同进行调整。
	kb.SetKeys(keybd_event.VK_VOLUME_DOWN)
	for {
		time.Sleep(1 * time.Second) // 可根据需要调整此间隔。
		err = kb.Launching()
		kb.Release()
		if err != nil {
			fmt.Println("Error simulating key press:", err)
		}
	}
}
