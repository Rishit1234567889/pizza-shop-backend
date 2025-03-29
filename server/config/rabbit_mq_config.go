package config

import (
	"fmt"
	"log"

	"strconv"

	"github.com/Rishit1234567889/pizza-shop/logger"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQConnection struct {
	conn  *amqp091.Connection
	queue string
}

func GetNewRabbitMQConnection() *RabbitMQConnection {
	host := GetEnvProperty("rabbit_mq_host")
	port := GetEnvProperty("rabbit_mq_port")
	username := GetEnvProperty("rabbit_mq_username")
	password := GetEnvProperty("rabbit_mq_password")
	queue := GetEnvProperty("rabbit_mq_default_queue")

	PORT, err := strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("invalid port: %v", err))
	}
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, PORT)
	conn, err := amqp091.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("failed to connect RabbitMQ: %v", err))
	}
	log.Printf(" RabbitMQ has been connected")
	return &RabbitMQConnection{
		conn:  conn,
		queue: queue,
	}
}

func (r *RabbitMQConnection) Connect() *amqp091.Connection {
	host := GetEnvProperty("rabbit_mq_host")
	port := GetEnvProperty("rabbit_mq_port")
	username := GetEnvProperty("rabbit_mq_username")
	password := GetEnvProperty("rabbit_mq_password")

	PORT, err := strconv.Atoi(port)
	if err != nil {
		panic(fmt.Sprintf("invalid port: %v", err))
	}
	url := fmt.Sprintf("amqp://%s:%s@%s:%d/", username, password, host, PORT)
	conn, err := amqp091.Dial(url)
	if err != nil {
		panic(fmt.Sprintf("failed to connect RabbitMQ: %v", err))
	}
	log.Printf(" RabbitMQ has been reconnected")
	return conn
}

func (r *RabbitMQConnection) DeclareQueue(queueName string) error {
	var err error
	channel, err := r.conn.Channel()
	defer channel.Close()

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (r *RabbitMQConnection) GetConnection() *amqp091.Connection {
	if r.conn == nil {
		r.conn = r.Connect()
	}
	return r.conn
}

func (r *RabbitMQConnection) GetChannel() *amqp091.Channel {
	var channel *amqp091.Channel
	connection := r.conn
	if connection == nil {
		connection = r.Connect()
	}
	channel, err := r.conn.Channel()
	if err != nil {
		channel, _ = r.conn.Channel()
		logger.Log("channel is nil")
	}

	if channel != nil && channel.IsClosed() {
		channel, err = r.conn.Channel()
		if err != nil {
			logger.Log("channel was closed and error creating channel")
		}
	}
	return channel
}

func (r *RabbitMQConnection) GetQueue() string {
	return r.queue
}

func (r *RabbitMQConnection) Close() {
	r.conn.Close()
}
