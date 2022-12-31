

import time
import pika


host = 'localhost'
port = 5672

exchange_name = "my-exchange"
queue_name = "my-queue"

# connect to RabbitMQ server and create a channel
connection_params = pika.ConnectionParameters(host=host, port=port)
connection = pika.BlockingConnection(connection_params)

channel = connection.channel()

# create an exhange
channel.exchange_declare(exchange=exchange_name, exchange_type='fanout')

# create a queue
queue = channel.queue_declare(queue=queue_name)

# bind the exchange and queue
channel.queue_bind(exchange=exchange_name, queue=queue_name)

print("Input messages to produce or `stop` to stop")

while True:
	message = input("Enter message to send or stop to exit: ")

	channel.basic_publish(exchange=exchange_name, routing_key='', body=message)
	print("message published")

	if message.lower() == "stop":
		break

# sleep for 2s waiting for consumer to receive stop message and delete queue
time.sleep(2)
channel.queue_delete(queue_name)
channel.exchange_delete(exchange_name)


connection.close()
