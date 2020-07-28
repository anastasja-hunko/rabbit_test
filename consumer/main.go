package main

import (
	"log"
	"time"

	"github.com/streadway/amqp"
)

const (
	AMQP_url   = "amqp://guest:guest@localhost:5672"
	QUEUE_name = "rabbit_test1"
)

func main() {
	//connect to RabbitMQ server
	conn, err := amqp.Dial(AMQP_url)

	printError(err, connectError)

	defer conn.Close()

	//create a channel
	ch, err := conn.Channel()

	printError(err, openChannelError)

	defer ch.Close()

	//declare the queue from which consuming messages
	_, err = ch.QueueDeclare(
		QUEUE_name,
		true,
		false,
		false,
		false,
		nil,
	)

	printError(err, declareQueueError)

	//read the messages from a channel
	msg, err := ch.Consume(
		QUEUE_name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	printError(err, registerConsumerError)

	forever := make(chan bool)

	go func() {
		for d := range msg {
			log.Printf("Received a message: %s", d.Body)

			time.Sleep(2 * time.Second)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
