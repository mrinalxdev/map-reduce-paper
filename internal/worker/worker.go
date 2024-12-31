package worker

import (
	"context"
	"encoding/json"
	"strings"

	// "encoding/json"

	"github.com/mrinalxdev/map-red/internal/models"
	"github.com/mrinalxdev/map-red/internal/queue"
	"github.com/mrinalxdev/map-red/internal/storage"
	"github.com/streadway/amqp"
)

type Worker struct {
	queue *queue.RabbitMQ
	storage *storage.Redis
}

func NewWorker(queue *queue.RabbitMQ, storage *storage.Redis) *Worker {
	return &Worker{
		queue : queue,
		storage: storage,
	}
}

func (w *Worker) processMapTask(ctx context.Context, delivery amqp.Delivery) {
    var task models.Task
    if err := json.Unmarshal(delivery.Body, &task); err != nil {
        delivery.Nack(false, true)
        return
    }

    // Example map function (word count)
    words := strings.Fields(string(task.Data))
    wordCount := make(map[string]int)

    for _, word := range words {
        wordCount[word]++
    }

    // Store results in Redis
    for word, count := range wordCount {
        if err := w.storage.StoreMapResult(ctx, word, string(count)); err != nil {
            delivery.Nack(false, true)
            return
        }
    }

    delivery.Ack(false)
}

func (w *Worker) processReduceTask(ctx context.Context, delivery amqp.Delivery) {
    var task models.Task
    if err := json.Unmarshal(delivery.Body, &task); err != nil {
        delivery.Nack(false, true)
        return
    }

    // Example reduce function (sum counts for a word)
    total := 0
    for _, value := range task.Values {
        count := 0
        if err := json.Unmarshal([]byte(value), &count); err != nil {
            delivery.Nack(false, true)
            return
        }
        total += count
    }

    // Store final result
    if err := w.storage.StoreMapResult(ctx, task.Key, string(total)); err != nil {
        delivery.Nack(false, true)
        return
    }

    delivery.Ack(false)
}

func (w *Worker) Run(ctx context.Context) error {
    mapCh, err := w.queue.Channel.Consume(
        "map_tasks",
        "",    // consumer
        false, 
        false, 
        false, 
        false, 
        nil,   
    )
    if err != nil {
        return err
    }

    reduceCh, err := w.queue.Channel.Consume(
        "reduce_tasks",
        "",    
        false, // auto-ack
        false, 
        false, // no-local
        false, // no-wait
        nil,  
    )
    if err != nil {
        return err
    }

    for {
        select {
        case <-ctx.Done():
            return ctx.Err()
        case delivery := <-mapCh:
            w.processMapTask(ctx, delivery)
        case delivery := <-reduceCh:
            w.processReduceTask(ctx, delivery)
        }
    }
}