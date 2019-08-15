package lib_kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

var (
	topic     = "hello"
	partition = int32(-1)
	addr      = []string{"localhost:9092"}
	key       = sarama.StringEncoder("key")
)

func NewKafkaClientProducer(addr []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	return sarama.NewSyncProducer(addr, config)
}

func producer() {
	addr := []string{"localhost:9092"}

	producer, err := newKafkaClientProducer(addr)
	if err != nil {
		panic(err)
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: partition,
		Key:       key,
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
