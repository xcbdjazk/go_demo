package main

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func consumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	// consumer
	consumer, err := sarama.NewConsumer(address, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	defer consumer.Close()
	if len(partitionList) == 0{
		return
	}

	partitionConsumer, err := consumer.ConsumePartition(topic, partitionList[0], sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partitionConsumer.Close()
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}
}
type Consumer struct {
	ready        chan bool
	consumerName string
}

func consumerGroup() {
	consumer := &Consumer{
		ready:        make(chan bool),
		consumerName: "consumer1",
	}
	consumer2 := &Consumer{
		ready:        make(chan bool),
		consumerName: "consumer2",
	}
	ctx, consumerCancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go startConsumer(ctx, []string{topic}, consumer)

	go startConsumer(ctx, []string{topic}, consumer2)
	wg.Wait()
	consumerCancel()
}


func (consumer *Consumer) Setup(_ sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}
func (consumer *Consumer) Cleanup(_ sarama.ConsumerGroupSession) error {
	fmt.Println("Cleanup")
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for message := range claim.Messages() {
		fmt.Printf("ConsumerName:%s Message claimed: value = %s, topic = %s,partition=%d \n", consumer.consumerName, string(message.Value), message.Topic, message.Partition)
		session.MarkMessage(message, "")
	}
	return nil
}

func startConsumer(ctx context.Context, topics []string, consumer *Consumer) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	//config.Version = sarama.V0_11_0_2
	//config.Net.SASL.Enable = true
	//config.Net.SASL.User = "admin"
	//config.Net.SASL.Password = "admin"
	// consumer
	consumerGroup, err := sarama.NewConsumerGroup(address, "test-group", config)
	if err != nil {
		fmt.Printf("NewConsumerGroup create consumer error %s\n", err.Error())
		return
	}
	go func() {
		for {
			if err := consumerGroup.Consume(context.Background(), topics, consumer); err != nil {
				fmt.Printf("Error from consumer: %v", err)
			}
			consumer.ready = make(chan bool)
		}
	}()
	<-consumer.ready
}