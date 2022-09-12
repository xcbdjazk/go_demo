package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"time"
)

type Config struct {
	App struct {
		Ip   string `json:"ip"`
		Port int    `json:"port"`
	} `json:"app"`
	DB struct {
		Mysql string `json:"mysql"`
		Redis string `json:"redis"`
	} `json:"db"`
}

func main() {
	// filePath := "config.json"
	filePath := "config.yaml"

	var conf Config

	_ = config.LoadFile(filePath)
	_ = config.Get().Scan(&conf)
	fmt.Print(conf)
	go func() {
		// 监听文件变更
		for {
			w, _ := config.Watch()
			v, _ := w.Next()
			v.Scan(&conf)
			fmt.Println("edit conf", conf)
		}
	}()
	time.Sleep(time.Second * 200)
}
