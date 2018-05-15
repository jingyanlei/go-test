package main

import (
	"log"
	"fmt"
	"github.com/Shopify/sarama"
	"gopkg.in/bsm/sarama-cluster.v1"
	"time"
	"strings"
)

var brokers = "127.0.0.1:9092"
var zks = "127.0.0.1:2181"
var group = "my-group4"
var topics = "my-topic4"

//var config cluster.Config

func main() {
	
	client, err := sarama.NewClient(strings.Split(brokers, ","), nil)
	abortOn(err)
	defer client.Close()
	//
	//err = produce(&client, 10000)
	//abortOn(err)
	//log.Println("Produced 10000 messages.")
	//
	err = consume(&client, 1000, func(message *sarama.ConsumerMessage) {
		fmt.Println(message.Topic, message.Partition, message.Offset, string(message.Value))
	})
	abortOn(err)
	log.Println("Consumed 1000 messages.")
	
	log.Println("Press ^C to exit.")
	select {}
	
	
	
}

func produce(client *sarama.Client, n int) error {
	//producer, err := sarama.NewSyncProducerFromClient(client)
	producer, err := sarama.NewSyncProducer(strings.Split(brokers, ","), nil)
	if err != nil {
		return err
	}
	defer producer.Close()

	for i := 0; i < n; i++ {
		kv := sarama.StringEncoder(fmt.Sprintf("MESSAGE-%08d", i))
		msg := sarama.ProducerMessage{
			Topic:topics,
			Key:kv,
			Value:kv,
		}
		
		if _, _, err := producer.SendMessage(&msg); err != nil {
			return err
		}
	}
	return nil
}

func consume(client *sarama.Client, n int, cb func(*sarama.ConsumerMessage)) error {
	// Connect consumer
	config := cluster.Config{
		CommitEvery: time.Second,
		DefaultOffsetMode: sarama.OffsetNewest,
		ZKSessionTimeout: time.Second * 30,
	}
	
	//consumer, err := cluster.NewConsumerFromClient(client, []string{"localhost:2181"}, "my-group", []string{"my-topic"}, &config)
	consumer, err := cluster.NewConsumer(strings.Split(brokers, ","), strings.Split(zks, ","), group, strings.Split(topics, ","), &config)
	if err != nil {
		return err
	}
	defer consumer.Close()

	// Consume and print out errors in a separate goroutine
	go func() {
		for msg := range consumer.Errors() {
			fmt.Println("ERROR:", msg.Error())
		}
	}()

	// Process and ack messages
	//counter := 0
	for msg := range consumer.Messages() {
		cb(msg)
		consumer.Ack(msg)
		//consumer.Commit()

		//if counter++; counter >= n {
		//	return nil
		//}
	}
	return nil
}

func abortOn(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
