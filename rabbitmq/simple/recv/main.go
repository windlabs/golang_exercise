package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func recv() {
	var (
		conn      *amqp.Connection
		err       error
		channel   *amqp.Channel
		queue     amqp.Queue
		deliveMsg <-chan amqp.Delivery
	)

	//建立连接
	if conn, err = amqp.Dial("amqp://liujq01:liujq01@192.168.33.10//test"); err != nil {
		log.Fatalf("%s:%s", "Failed to connect rabbitmq", err)
		return
	}
	defer conn.Close()

	//创建channel
	if channel, err = conn.Channel(); err != nil {
		log.Fatalf("%s:%s", "channel create failed.", err)
		return
	}
	defer channel.Close()

	if queue, err = channel.QueueDeclare(
		"test",
		true,
		false,
		false,
		false,
		nil); err != nil {
		log.Fatalf("%s:%s", "declare queue failed.", err)
		return
	}

	if deliveMsg, err = channel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	); err != nil {
		log.Fatalf("%s:%s", "consume msg failed.", err)
		return
	}
	go func() {
		for d := range deliveMsg {
			fmt.Printf("Recv:%s", d.Body)
			d.Ack(false)
		}
	}()

	for {

	}
}

func main() {
	recv()
}
