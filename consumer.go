package main

import (
	"github.com/streadway/amqp"
	"time"
	"log"
)

func CreateConsumer(rabbitMqIp string, rabbitMQPort int, queueName string, channel chan *amqp.Delivery) {
	for {
		consume(rabbitMqIp, rabbitMQPort, queueName, channel)
		time.Sleep(FREQUENCY_CHECK_QUEUES)
	}
}

func consume(rabbitMqIp string, rabbitMQPort int, queueName string, channel chan *amqp.Delivery) {
	msgs, err := InitializeAndGetMessages(rabbitMqIp, rabbitMQPort, queueName)

	if err!=nil {
		log.Println("error connecting to queue " + err.Error())
		return
	}

	for d := range msgs {
		channel <- &d
	}
}
