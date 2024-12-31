package main

import (
	"context"
	"log"

	"github.com/mrinalxdev/map-red/internal/config"
	"github.com/mrinalxdev/map-red/internal/master"
	"github.com/mrinalxdev/map-red/internal/queue"
	"github.com/mrinalxdev/map-red/internal/storage"
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

    master := master.NewMaster(queue, storage)
    
    ctx := context.Background()
    if err := master.Run(ctx, []byte("input data")); err != nil {
        log.Fatal(err)
    }
}