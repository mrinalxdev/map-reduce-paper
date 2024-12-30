package config

import "os"

type Config struct {
	RedisURL string
	RabbitMQURL string
}

func LoadConfig() (*Config, error){
	return &Config {
		RedisURL: getEnvOrDefualt("REDIS_URL", "localhost:6397"),
		RabbitMQURL: getEnvOrDefualt("RABBITMQ_URL", "amqp://guest@localhost:5672/"),
	}, nil
}

func getEnvOrDefualt(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}