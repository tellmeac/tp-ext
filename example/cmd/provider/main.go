// This application provides messages to the queue with periodicity
package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/nats-io/nats.go"
)

type Config struct {
	NatsURL string `env:"NATS_URL"`
}

func main() {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Config setup: %s", err)
	}

	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatalf("Nats connection: %s", err)
	}

	data, _ := json.Marshal(map[string]string{
		"message": "hello world",
	})

	ticker := time.NewTicker(time.Minute)
	for {
		<-ticker.C
		nc.Publish("test.*", data)
	}
}
