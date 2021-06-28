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
// MyHandler Class
type MyHandler struct {
	Title string
}

func (m *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v, msg:%v\n", m.Title, msg.NSQDAddress, string(msg.Body))
	return
}

func initConsumer(topic string, channel string, addr string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	c, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Printf("create consumer failed, err:%v\n", err)
		return
	}
	consumer := &MyHandler{
		Title: "title",
	}
	c.AddHandler(consumer)

	//if err := c.ConnectToNSQD(addr); err != nil {
	if err := c.ConnectToNSQLookupd(addr); err != nil {
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
	c := make(chan os.Signal)        //定义一个信号的通道
	signal.Notify(c, syscall.SIGINT) //转发键盘中断信号到通道
	<-c                              //阻塞
}
