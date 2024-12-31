package queue

import (
	"encoding/json"

	"github.com/mrinalxdev/map-red/internal/models"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	Channel *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error){
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	for _, queue := range []string{"map_tasks", "reduce_tasks"}{
		_, err = ch.QueueDeclare(
			queue,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			return nil, err
		}
	}

	return &RabbitMQ{
		conn : conn,
		Channel: ch,
	}, nil
}

func (r *RabbitMQ) PublishTask(queueName string, task *models.Task) error {
    body, err := json.Marshal(task)
    if err != nil {
        return err
    }

    return r.Channel.Publish(
        "",       
        queueName, 
        false,     
        false,     
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  "application/json",
            Body:        body,
        },
    )
}