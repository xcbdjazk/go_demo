package main

var (
	address = []string{"192.168.2.170:9092", "192.168.2.170:9093", "192.168.2.170:9094"}
	topic   = "first"
)

func main() {

	go producer()
	go consumer()
	go consumerGroup()
	select {}
	//fmt.Println( pid, offset)

}
