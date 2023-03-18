package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"net/http"
	"time"
)

func producer() {

	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Partitioner = sarama.NewRandomPartitioner
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true
	cfg.Producer.Timeout = time.Second

	cli, err := sarama.NewAsyncProducer(address, cfg)
	if err != nil {
		fmt.Println(err)
		return
	}
	go func() {
		for err := range cli.Errors() {
			fmt.Println(err)
		}

	}()
	go func() {
		for i := range cli.Successes() {
			_ = i
		}

	}()

	defer cli.Close()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method)
		msg := request.URL.Query().Get("msg")
		kmsg := &sarama.ProducerMessage{}
		kmsg.Topic = "first"
		kmsg.Value = sarama.StringEncoder(msg)
		cli.Input() <- kmsg
		_, _ = writer.Write([]byte("ok"))
	})
	fmt.Println("ListenAndServe :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}

}
