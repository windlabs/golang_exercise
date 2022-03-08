package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	var (
		conn        *amqp.Connection
		err         error
		chanel      *amqp.Channel
		queue       amqp.Queue
		deliveryMsg <-chan amqp.Delivery
	)
	// 消费者同样需要建立连接和channel、然后申明我们想消费的chanel
	// 连接amqp-server
	if conn, err = amqp.Dial("amqp://bairimeng:123123@140.143.249.189:5672//test"); err != nil {
		failOnError(err, "Failed to connect to RabbitMQ")
		return
	}
	defer conn.Close()
	// 创建channel
	if chanel, err = conn.Channel(); err != nil {
		failOnError(err, "Failed to create channel")
		return
	}
	defer chanel.Close()

	err = chanel.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		false,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	// 为了发送消息，我们必须声明一个queue，让后往这个queue中发送消息
	if queue, err = chanel.QueueDeclare(
		"",    // todo name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	); err != nil {
		failOnError(err, "Failed to declare queue")
		return
	}
	// 将queue和交换机绑定
	err = chanel.QueueBind(
		queue.Name, // queue name
		"",         // routing key
		"logs",     // exchange
		false,
		nil,
	)
	failOnError(err, "Failed to bind a queue")

	// 接受消息,函数返回值是chan，当消费者接收到消息之后，我们可从 deliveryMsg 中将消息取出
	if deliveryMsg, err = chanel.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack 默认自动ACK（有丢失消息的风险）
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	); err != nil {
		failOnError(err, "Failed to Consume msg")
		return
	}

	// 使用消费空队列的方式：防止程序的自动推出
	forever := make(chan bool)
	// 启动单独的协程消费chan
	go func() {
		for d := range deliveryMsg {
			fmt.Printf("Recv: %s", d.Body)
			d.Ack(false)
		}
	}()
	<-forever
}

// 处理amqp调用失败时产生的错误
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
