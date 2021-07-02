package consumer

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (
	topic      string
	channel    string
	lookupAddr string
	consumer   *nsq.Consumer
)

func init() {
	initConfig()

	topic = viper.GetString("nsq.topic")
	channel = viper.GetString("nsq.channel")
	lookupAddr = fmt.Sprintf("%s:%d", viper.GetString("nsq.lookup.host"), viper.GetInt("nsq.lookup.port"))

	initConsumer()
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.AddConfigPath("conf")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../conf")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("../../conf")
	viper.SetConfigName("app")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("init config failed, err:%v\n", err)
		return err
	}
	return nil
}

type Handler struct {
	task func(addr, msg string) error
}

func (h *Handler) HandleMessage(msg *nsq.Message) error {
	err := h.task(msg.NSQDAddress, string(msg.Body))
	if err != nil {
		log.Printf("task handle failed, err:%v\n", err)
		return err
	}
	//log.Printf("from %v, msg:%v\n", msg.NSQDAddress, string(msg.Body))
	return nil
}

func initConsumer() error {
	var err error
	conf := nsq.NewConfig()
	conf.LookupdPollInterval = 1 * time.Minute
	consumer, err = nsq.NewConsumer(topic, channel, conf)
	if err != nil {
		log.Printf("init consumer failed, err:%v\n", err)
		return err
	}
	return nil
}

func Consumer(task func(addr, msg string) error) error {
	handler := new(Handler)
	handler.task = task
	consumer.AddHandler(handler)

	//if err := consumer.ConnectToNSQD(lookupAddr); err != nil {
	if err := consumer.ConnectToNSQLookupd(lookupAddr); err != nil {
		log.Printf("subscribe failed, err:%v\n", err)
		return err
	}
	return nil
}
