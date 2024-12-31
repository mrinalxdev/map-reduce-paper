package main

import (
	"context"
	"log"

	"github.com/mrinalxdev/map-red/internal/config"
	"github.com/mrinalxdev/map-red/internal/queue"
	"github.com/mrinalxdev/map-red/internal/storage"
	"github.com/mrinalxdev/map-red/internal/worker"
)

func main() {
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatal(err)
    }

    queue, err := queue.NewRabbitMQ(cfg.RabbitMQURL)
    if err != nil {
        log.Fatal(err)
    }

    storage, err := storage.NewRedis(cfg.RedisURL)
    if err != nil {
        log.Fatal(err)
    }

    worker := worker.NewWorker(queue, storage)
    
    ctx := context.Background()
    if err := worker.Run(ctx); err != nil {
        log.Fatal(err)
    }
}