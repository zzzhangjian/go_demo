package ctl_kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/zzzhangjian/go_demo/library/kafka"
)

var (
	topic     = "hello"
	partition = int32(-1)
)

func producer() {

	addr := []string{"localhost:9092"}

	producer, err := NewKafkaClientProducer(addr, config)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: int32(-1),
		Key:       sarama.StringEncoder("key"),
	}

	var value string
	for {
		_, err := fmt.Scanf("%s", &value)
		if err != nil {
			break
		}
		msg.Value = sarama.ByteEncoder(value)
		fmt.Println(value)

		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println("Send message Fail")
		}
		fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	}
}
