// This application provides messages to the queue with periodicity
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/nats-io/nats.go"
)

type Config struct {
	NatsURL       string `env:"NATS_URL"`
	ListenAddress string `env:"LISTEN_ADDRESS" env-default:"0.0.0.0:80"`
}

var acceptedCount uint64 = 0

func main() {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Config setup: %s", err)
	}

	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		log.Fatalf("Nats connection: %s", err)
	}

	nc.Subscribe("test.*", func(msg *nats.Msg) {
		acceptedCount++
		log.Printf("Accepted another message. Total accepted count: %d", acceptedCount)
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		data, _ := json.Marshal(map[string]interface{}{
			"count": acceptedCount,
		})

		w.Write(data)
	})

	if err := http.ListenAndServe(cfg.ListenAddress, mux); err != nil {
		log.Fatalf("Failed to listen http address: %s", err)
	}
}
