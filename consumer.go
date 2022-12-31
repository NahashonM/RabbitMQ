package main

import (
	"fmt"
	"strconv"
	"strings"

	amqp "github.com/rabbitmq/amqp091-go"
)

var username string = "guest"
var password string = "guest"
var host string = "localhost"
var port int = 5672

var queue_name string = "my-queue"

func main() {

	// connect to rabbitMQ
	connect_string := "amqp://" + username + ":" + password + "@" + host + ":" + strconv.Itoa(port)
	connection, _ := amqp.Dial(connect_string)

	// open a channel
	channel, _ := connection.Channel()

	msgs, _ := channel.Consume(
		queue_name, // queue string
		"",         // consumer string
		true,       // autoack
		false,      // exclusive
		false,      // no local
		false,      // no wait
		nil,        // args
	)

	forever := make(chan string)

	// go routine to handle messages
	go func(c chan string) {

		for d := range msgs {

			message := string(d.Body)

			fmt.Printf("msg: %s", message)

			if strings.HasPrefix(message, "stop") {
				break
			}
		}

		channel.Cancel("", false)
		connection.Close()

		c <- "done"

	}(forever)

	fmt.Println("Waiting for messages, send stop message to exit")

	// exit only when routine is done
	<-forever
}
