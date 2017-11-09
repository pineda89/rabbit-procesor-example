package main

import (
	"github.com/streadway/amqp"
	"strconv"
)

func InitializeAndGetMessages(rabbitIp string, rabbitPort int, rabbitQueue string) (<-chan amqp.Delivery, error) {

	delivery, err := rabbitInitializeAndGetMessages(rabbitIp, rabbitPort, rabbitQueue)

	return delivery, err
}

func rabbitInitializeAndGetMessages(rabbitIp string, rabbitPort int, rabbitQueue string) (<-chan amqp.Delivery, error) {
	connection, err := amqp.Dial("amqp://" + rabbitIp + ":" + strconv.Itoa(rabbitPort) + "/")
	if err!=nil {
		return nil, err
	}

	ch, err := connection.Channel()
	if err!=nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		rabbitQueue, // name
		true,   // durable
		false,   // delete when usused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err!=nil {
		return nil, err
	}
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err!=nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	return msgs, err
}