package kafka

import (
	"fmt"
	"github.com/murat96k/kitaptar.kz/internal/user/config"
	"log"

	"github.com/IBM/sarama"

	"strings"
)

type ConsumerCallback interface {
	Callback(message <-chan *sarama.ConsumerMessage, error <-chan *sarama.ConsumerError)
}

type Consumer struct {
	topics   []string
	master   sarama.Consumer
	callback ConsumerCallback
}

func NewConsumer(
	cfg config.Kafka,
	callback ConsumerCallback,
) (*Consumer, error) {
	samaraCfg := sarama.NewConfig()
	samaraCfg.ClientID = "go-kafka-consumer"
	samaraCfg.Consumer.Return.Errors = true

	master, err := sarama.NewConsumer(cfg.Brokers, samaraCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create NewConsumer err: %w", err)
	}

	return &Consumer{
		topics:   cfg.Consumer.Topics,
		master:   master,
		callback: callback,
	}, nil
}

func (c *Consumer) Start() {
	consumers := make(chan *sarama.ConsumerMessage)
	errors := make(chan *sarama.ConsumerError)

	for _, topic := range c.topics {
		if strings.Contains(topic, "__consumer_offsets") {
			continue
		}

		partitions, _ := c.master.Partitions(topic)

		consumer, err := c.master.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
		if nil != err {
			log.Printf("Topic %v Partitions: %v, err: %v", topic, partitions, err)
			continue
		}
		log.Printf(" Start consuming topic %v", topic)

		go func(topic string, consumer sarama.PartitionConsumer) {
			for {
				select {
				case consumerError := <-consumer.Errors():
					errors <- consumerError

				case msg := <-consumer.Messages():
					consumers <- msg
				}
			}
		}(topic, consumer)
	}

	c.callback.Callback(consumers, errors)
}
