package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var host string = "localhost"
var port int = 5672

var queue_name string = "my-queue"
var exchange_name string = "my-exchange"

func main() {

	// connect to rabbitMQ
	conn, _ := amqp.Dial("amqp://guest:guest@localhost:5672/")

	// open a channel
	ch, _ := conn.Channel()

	// declare queue
	ch.QueueDeclare(
		queue_name, // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)

	// declare an exchange
	ch.ExchangeDeclare(
		exchange_name, // name
		"fanout",      // exchange type
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)

	ch.QueueBind(
		queue_name,    // destination
		"",            // key
		exchange_name, // source
		false,         // nowait
		nil,           // args
	)

	reader := bufio.NewReader(os.Stdin)

	for true {

		fmt.Print("Enter message to publish or stop exit: ")
		message, _ := reader.ReadString('\n')

		m := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		}

		ch.Publish(
			exchange_name, // exchange name
			queue_name,    // key
			false,         // mandatory
			false,         // immediate
			m,             // message body
		)

		if strings.HasPrefix(message, "stop") {
			break
		}

		time.Sleep(time.Millisecond * 10)
	}

	// delete queue
	ch.QueueDelete(
		queue_name, // queue name
		false,      // ifunused
		false,      // ifempty
		false,      // nowait
	)

	// delete channel
	ch.ExchangeDelete(
		exchange_name, // exchange name
		false,         // ifunused
		false,         // nowait
	)

	// close connection
	defer conn.Close()
}
