package main

import (
	"github.com/streadway/amqp"
)

func CreateWorker(channel chan *amqp.Delivery) {

	for {
		d := <- channel

		doSomethingWithQueueMessage(d.Body)
	}

}

func doSomethingWithQueueMessage(msg []byte) {

}

