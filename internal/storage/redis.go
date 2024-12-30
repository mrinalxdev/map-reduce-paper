package storage

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(url string) (*Redis, error){
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	client := redis.NewClient(opt)
	return &Redis{client : client}, nil
}

func (r *Redis) StoreMapResult(ctx context.Context, key, value string) error{
	return r.client.HSet(ctx, "map_results", key, value).Err()
}

func (r *Redis) GetMapResults(ctx context.Context) (map[string]string, error){
	return r.client.HGetAll(ctx, "map_results").Result()
}