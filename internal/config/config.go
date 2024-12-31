
package config

import (
    "fmt"
    "os"
)

type Config struct {
    RedisURL    string
    RabbitMQURL string
}

func LoadConfig() (*Config, error) {
    // Default to local development URLs
    redisHost := getEnvOrDefault("REDIS_HOST", "localhost")
    redisPort := getEnvOrDefault("REDIS_PORT", "6379")
    rabbitmqHost := getEnvOrDefault("RABBITMQ_HOST", "localhost")
    rabbitmqPort := getEnvOrDefault("RABBITMQ_PORT", "5672")

    return &Config{
        RedisURL:    fmt.Sprintf("redis://%s:%s", redisHost, redisPort),
        RabbitMQURL: fmt.Sprintf("amqp://guest:guest@%s:%s/", rabbitmqHost, rabbitmqPort),
    }, nil
}

func getEnvOrDefault(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}