package lib_kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

var (
	wg sync.WaitGroup
)

func consumer() {
	topic = "hello"
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}

	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		defer pc.AsyncClose()

		wg.Add(1)

		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
