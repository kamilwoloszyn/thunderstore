package config

type Config struct {
	EnvKafkaHost string `env:"ENV_KAFKA_HOST" envDefault:""`
	EnvKafkaPort string `env:"ENV_KAFKA_PORT" envDefault:""`
}
