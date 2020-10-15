package main

import (
	"log"
	"net/http"
	"os"
)

func SetupLogger() {
	logFileLocation, _ := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	log.SetOutput(logFileLocation)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url %s : %s", url, err.Error())
	} else {
		log.Printf("Status Code for %s : %s", url, resp.Status)
		defer resp.Body.Close()
	}
}

//func main() {
//	SetupLogger()
//	simpleHttpGet("www.baidu.com")
//	simpleHttpGet("http://www.baidu.com")
//}
func main() {
	SetupLogger()
	Ldefault()
	Ldate()
	Ltime()
	Lmicroseconds()
	Llongfile()
	Lshortfile()
	LUTC()
}

func Ldefault() {
	log.Println("这是默认的格式\n")
}

func Ldate() {
	log.SetFlags(log.Ldate)
	log.Println("这是输出日期格式\n")
}

func Ltime() {
	log.SetFlags(log.Ltime)
	log.Println("这是输出时间格式\n")
}

func Lmicroseconds() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("这是输出微秒格式\n")
}

func Llongfile() {
	log.SetFlags(log.Llongfile)
	log.Println("这是输出路径+文件名+行号格式\n")
}

func Lshortfile() {
	log.SetFlags(log.Lshortfile)
	log.Println("这是输出文件名+行号格式\n")
}

func LUTC() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.Println("这是输出 使用标准的UTC时间格式 格式\n")
}
