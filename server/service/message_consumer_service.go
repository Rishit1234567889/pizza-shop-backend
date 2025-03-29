package service

import (
	"fmt"

	"github.com/Rishit1234567889/pizza-shop/config"
	"github.com/Rishit1234567889/pizza-shop/logger"
	"github.com/rabbitmq/amqp091-go"
)

type IMessageConsumerService interface {
	DeclareQueue(queueName string) error
	ConsumeEventAndProcess(queueName string, processor IMessageProcessor) error
}

type MessageConsumerService struct {
	conf *config.RabbitMQConnection
}

func (mcs *MessageConsumerService) DeclareQueue(queueName string) error {
	channel := mcs.conf.GetChannel()
	if channel == nil {
		return fmt.Errorf("message channel is nil, please retry")
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

func (mcs *MessageConsumerService) ConsumeEventAndProcess(queueName string, processor IMessageProcessor) error {

	channel := mcs.conf.GetChannel()
	if channel == nil {
		return fmt.Errorf("messaging channel is nil, retry")
	}

	logger.Log("consuming message....")
	msgs, err := channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return fmt.Errorf("failed to consume message: %w", err)
	}

	go func() {
		for msg := range msgs {
			go func(msg amqp091.Delivery) {
				err := processor.ProcessMessage(msg)
				if err != nil {
					logger.Log(fmt.Sprintf("Message processing failed: %v", err))
				}
			}(msg)
		}
	}()

	select {}

}

func GetMessageConsumerService() *MessageConsumerService {
	rabbitMQConf := config.GetNewRabbitMQConnection()
	return &MessageConsumerService{
		conf: rabbitMQConf,
	}
}
