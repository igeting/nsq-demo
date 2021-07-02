package producer

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
)
import "github.com/spf13/viper"

var (
	topic    string
	channel  string
	nsqdAddr string
	producer *nsq.Producer
)

func init() {
	initConfig()

	topic = viper.GetString("nsq.topic")
	channel = viper.GetString("nsq.channel")
	nsqdAddr = fmt.Sprintf("%s:%d", viper.GetString("nsq.nsqd.host"), viper.GetInt("nsq.nsqd.port"))

	initProducer()
}

func initConfig() error {
	viper.AddConfigPath(".")
	viper.AddConfigPath("configs")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("app")

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("init config failed, err:%v\n", err)
		return err
	}
	return nil
}

func initProducer() error {
	var err error
	conf := nsq.NewConfig()
	producer, err = nsq.NewProducer(nsqdAddr, conf)
	if err != nil {
		log.Printf("init producer failed, err:%v\n", err)
		return err
	}
	return nil
}

func Producer(msg string) error {
	err := producer.Publish(topic, []byte(msg))
	if err != nil {
		log.Printf("publish failed, err:%v\n", err)
		return err
	}
	return nil
}
