/**
 * @author jiangshangfang
 * @date 2022/2/23 8:11 PM
 **/

package main

import (
	"fmt"
	"gin-micro-mine/pkg/queue/rabbitmq"
	"log"
)

func main() {
	addr := "root:12345678@localhost:5672/test"
	exchangeName := "test-exchange"
	queueName := "test-bind-to-exchange"
	routingKey := "test-route-key"
	var message = "Hello World RabbitMQ!"

	producer := rabbitmq.NewProducer(addr, exchangeName)
	defer producer.Stop()
	if err := producer.Start(); err != nil {
		log.Fatalf("start producer err: %s", err.Error())
	}
	if err := producer.Publish(routingKey, message); err != nil {
		log.Fatalf("failed publish message: %s", err.Error())
	}

	handler := func(body []byte) error {
		fmt.Println("consumer handler receive msg: ", string(body))
		return nil
	}

	consumer := rabbitmq.NewConsumer(addr, exchangeName, queueName, false, handler)
	if err := consumer.Start(); err != nil {
		log.Fatalf("failed consume: %s", err)
	}
}
