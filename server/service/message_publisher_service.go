package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Rishit1234567889/pizza-shop/config"
	"github.com/Rishit1234567889/pizza-shop/logger"
	"github.com/rabbitmq/amqp091-go"
)

type IMessagePublisher interface {
	PublishEvent(queueName string, body interface{}) error
	DeclareQueue(queueName string) error
}

type MessagePublisher struct {
	conf *config.RabbitMQConnection
}

func (mp *MessagePublisher) DeclareQueue(queueName string) error {
	channel := mp.conf.GetChannel()
	if channel == nil {
		return fmt.Errorf("message channel is nil, try again")
	}

	_, err := channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}

func (mp *MessagePublisher) PublishEvent(queueName string, body interface{}) error {
	data, err := json.Marshal(body)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if queueName == "" {
		queueName = config.GetEnvProperty("rabbit_mq_default_queue")
	}
	channel := mp.conf.GetChannel()
	if channel == nil {
		panic("message channel is nil, try again")
	}
	if channel.IsClosed() {
		panic("could not public channel, channel closed")
	}

	logger.Log(fmt.Sprintf("created new channel:%v", &channel))
	err = channel.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp091.Persistent,
		},
	)
	if err != nil {
		return err
	}
	logger.Log(fmt.Sprintf("Event published: %v", body))
	channel.Close()
	logger.Log(fmt.Sprintf("channel closed: %v", &channel))
	return nil

}

func GetMessagePublisherServer() *MessagePublisher {
	rabbitMQConf := config.GetNewRabbitMQConnection()
	return &MessagePublisher{
		conf: rabbitMQConf,
	}
}
