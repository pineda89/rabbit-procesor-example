package main

import (
	"github.com/streadway/amqp"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	QUEUE_ADDRESS string
	QUEUE_PORT int
	QUEUE_NAME string

	NUM_OF_CONSUMERS = 1
	NUM_OF_WORKERS = 10

	FREQUENCY_CHECK_QUEUES time.Duration = 1 * time.Second
)

func main() {
	consumeQueueAndProcess()

	// Block forever
	captureInterruptSignal()
}

func consumeQueueAndProcess() {
	channel := make(chan *amqp.Delivery, NUM_OF_WORKERS * 2)

	for i:=0; i<NUM_OF_CONSUMERS; i++ {
		go CreateConsumer(QUEUE_ADDRESS, QUEUE_PORT, QUEUE_NAME, channel)
	}
	for i:=0; i<NUM_OF_WORKERS; i++ {
		go CreateWorker(channel)
	}
}

// Function that blocks forever.
func captureInterruptSignal() {

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

}