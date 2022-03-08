package main

import (
	"github.com/streadway/amqp"
	"log"
)

func send() {
	var (
		conn    *amqp.Connection
		err     error
		channel *amqp.Channel
		queue   amqp.Queue
	)
	//连接amqp-server
	if conn, err = amqp.Dial("amqp://liujq01:liujq01@192.168.33.10//test"); err != nil {
		log.Fatalf("%s:%s", "Failed to connect to RabbitMQ", err)
		return
	}
	defer conn.Close()

	//创建channel
	if channel, err = conn.Channel(); err != nil {
		log.Fatalf("%s:%s", "Failed to create channel", err)
		return
	}
	defer channel.Close()

	//声明queue
	if queue, err = channel.QueueDeclare(
		"test",
		true,
		false,
		false,
		false,
		nil); err != nil {
		log.Fatalf("%s:%s", "Failed to create queue", err)
	}

	body := "Hello world2.\n"
	if err = channel.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		}); err != nil {
		log.Fatalf("%s:%s", "Failed to send msg", err)
		return
	}
}
func main() {
	send()
}
