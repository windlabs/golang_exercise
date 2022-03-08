package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	var (
		conn   *amqp.Connection
		err    error
		chanel *amqp.Channel
	)
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
	// 为channel绑定交换机
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

	// 发送消息
	body := "Hi bai ri meng"
	err = chanel.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to send msg")

}

// 处理amqp调用失败时产生的错误
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
	}
}
