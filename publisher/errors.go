package main

import "log"

const (
	connectError        = "Failed to connect to RabbitMQ"
	openChannelError    = "Failed to open a channel"
	declareQueueError   = "Failed to declare a queue"
	publishMessageError = "Failed to publish a message"
)

func printError(err error, message string) {

	if err != nil {
		log.Printf("%v : %v", message, err.Error())
	}

}
