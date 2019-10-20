package config

// Consumer holds the configuration values for the Kafka consumer.
type Consumer struct {
	Brokers []string `env:"KAFKA_BROKERS" default:"[192.168.178.26:9092]"`
	Topic   string   `env:"KAFKA_STORE_MUNICIPALITY_SUPERHERO_TOPIC" default:"update.municipality.superhero"`
	GroupID string   `env:"KAFKA_CONSUMER_REGISTER_GROUP_ID" default:"consumer-update-group"`
}