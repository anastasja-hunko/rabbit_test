package main

import "log"

const (
	connectError          = "Failed to connect to RabbitMQ"
	openChannelError      = "Failed to open a channel"
	declareQueueError     = "Failed to declare a queue"
	registerConsumerError = "Failed to register a consumer"
)

func printError(err error, message string) {

	if err != nil {
		log.Printf("%v : %v", message, err.Error())
	}

}
