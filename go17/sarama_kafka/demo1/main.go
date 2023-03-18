package main

var (
	address = []string{"192.168.163.121:9092"}
	topic   = "first"
)

func main() {

	go producer()
	go consumer()
	go consumerGroup()
	select {}
	//fmt.Println( pid, offset)

}
