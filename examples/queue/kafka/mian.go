/**
 * @author jiangshangfang
 * @date 2022/2/24 11:26 PM
 **/
package main

import (
	"github.com/Shopify/sarama"
	"github.com/jonny-mark/gin-micro-mine/pkg/queue/kafka"
	"log"
	"os"
)

func main() {
	var (
		config = sarama.NewConfig()
		logger = log.New(os.Stderr, "[sarama_logger]", log.LstdFlags)
		//groupID = "sarama_consumer"
		topic   = "go-message-broker-topic"
		brokers = []string{"localhost:9093"}
		message = "Hello World Kafka!"
	)

	// kafka publish message
	kafka.NewProducer(config, logger, topic, brokers).Publish(message)

	// kafka consume message
	//kafka.NewConsumer(config, logger, topic, groupID, brokers).Consume()

}
