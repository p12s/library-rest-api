package queueClient

import (
	"encoding/json"
	"fmt"

	"github.com/p12s/library-rest-api/library/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// QueueClient
type QueueClient struct {
	channel *amqp.Channel
	queue   *amqp.Queue
}

// New
func New(cfg config.Config) (*QueueClient, error) {
	connectionStr := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		cfg.Queue.User, cfg.Queue.Password, cfg.Queue.Host, cfg.Queue.Port)

	conn, err := amqp.Dial(connectionStr)
	if err != nil {
		logrus.Fatalf("error rabbitmq connect: %s\n", err.Error())
		return nil, fmt.Errorf("failed to connect to rabbitmq %w/n", err)
	}

	channel, err := conn.Channel()
	if err != nil {
		logrus.Fatalf("failed to open rabbitmq channel: %s\n", err.Error())
		return nil, fmt.Errorf("failed to open rabbitmq channel: %w/n", err)
	}

	queue, err := channel.QueueDeclare(
		cfg.Queue.Topic, // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		logrus.Fatalf("failed to declare rabbitmq: %s\n", err.Error())
		return nil, fmt.Errorf("failed to declare rabbitmq: %w/n", err)
	}

	return &QueueClient{
		channel: channel,
		queue:   &queue,
	}, nil
}

// Close - channel
func (q *QueueClient) Close() error {
	return q.channel.Close()
}

// Produce - send message to queue
func (q *QueueClient) Produce(payload interface{}) error {
	jsonString, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload to json string: %w", err)
	}

	err = q.channel.Publish(
		"",           // exchange
		q.queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{ContentType: "text/plain", Body: jsonString})

	if err != nil {
		return fmt.Errorf("rabbitmq channel publish: %w", err)
	}
	return nil
}
