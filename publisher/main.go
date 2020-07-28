package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

const (
	AMQP_url        = "amqp://guest:guest@localhost:5672"
	QUEUE_name      = "rabbit_test1"
	DEFAULT_message = "no_message"
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

	//declare a queue to send to
	_, err = ch.QueueDeclare(
		QUEUE_name,
		true,
		false,
		false,
		false,
		nil,
	)

	printError(err, declareQueueError)

	//infinity loop for publishing any message entered at console
	for {
		var message = DEFAULT_message
		fmt.Println("Enter a message:")
		fmt.Scanf("%v\n", &message)

		err = ch.Publish(
			"",
			QUEUE_name,
			false,
			false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			})

		printError(err, publishMessageError)
	}
}
