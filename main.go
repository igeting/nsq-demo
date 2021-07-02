package main

import (
	"log"
	"nsq-demo/api/consumer"
	"nsq-demo/api/producer"
	"time"
)

func main() {
	//consumer
	consumer.Consumer(
		//handler
		func(addr, msg string) error {
			log.Printf("from %v, msg:%v\n", addr, msg)
			return nil
		},
	)

	//producer
	for {
		producer.Producer("hello world")
		time.Sleep(time.Second * 1)
	}
}
