


import pika


host = 'localhost'
port = 5672

exchange_name = "my-exchange"
queue_name = "my-queue"

# connect to RabbitMQ server and create a channel
connection_params = pika.ConnectionParameters(host=host, port=port)
connection = pika.BlockingConnection(connection_params)

channel = connection.channel()

# define callback for whenever a message is available in the queue we are about to subscribe to
def callback(ch, method, properties, body):

	message = body.decode()

	if message.lower() == "stop":
		channel.stop_consuming()
		channel.close()
	
	print(message)


# consume messages sent to queue
print("Waiting for messages. Send `stop` or press `ctrl + c` to exit")

channel.basic_consume(queue=queue_name, on_message_callback=callback, auto_ack=True)
channel.start_consuming()