package main

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// NSQ Consumer Demo
type Handler struct {
	Title string
}

func (m *Handler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return nil
}

func initConsumer(topic string, channel string, addr string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	handler := &Handler{
		Title: "title",
	}
	consumer.AddHandler(handler)

	//if err := consumer.ConnectToNSQD(addr); err != nil {
	if err := consumer.ConnectToNSQLookupd(addr); err != nil {
		return err
	}
	return nil

}

func main() {
	topicAddress := "127.0.0.1:4161"
	err := initConsumer("topic_demo", "first", topicAddress)
	if err != nil {
		fmt.Printf("init consumer failed, err:%v\n", err)
		return
	}

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT)
	<-c
}
