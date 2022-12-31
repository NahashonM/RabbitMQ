# RabbitMQ - Hello-World [ Go ]

Note: The queue and exchange are defined on the producer <br>
       The producer should be run started and then the consumer can be started
## Producer

1. connect to RabbitMQ server and create a channel
1. define an exchange
1. create a queue
1. bind the exchange and queue
1. Publish messages until a stop is issued
1. Cleanup and delete the exchange and queue
1. Close the connection



## Consumer

1. connect to RabbitMQ server and create a channel
1. define a callback for messages on the queue
1. start consuming messages on the queue
1. Close the connection when stop is issued

