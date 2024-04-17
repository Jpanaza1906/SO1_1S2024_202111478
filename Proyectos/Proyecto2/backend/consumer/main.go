package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	// Configuración de Kafka
	kafkaBootstrapURL := "localhost:9092"
	kafkaTopic := "topic-sopes1"
	groupID := "my-consumer-group"
	autoOffsetReset := "earliest"

	config := kafka.ConfigMap{
		"bootstrap.servers": kafkaBootstrapURL,
		"group.id":          groupID,
		"auto.offset.reset": autoOffsetReset,
	}

	// Crear consumidor de Kafka
	c, err := kafka.NewConsumer(&config)
	if err != nil {
		fmt.Printf("Error al crear el consumidor de Kafka: %v\n", err)
		os.Exit(1)
	}

	// Suscribirse al topic
	err = c.SubscribeTopics([]string{kafkaTopic}, nil)
	if err != nil {
		fmt.Printf("Error al suscribirse al topic: %v\n", err)
		os.Exit(1)
	}

	// Manejar señales para cerrar el consumidor
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("Señal recibida: %v\n", sig)
			run = false
		default:
			// Leer mensaje del topic
			msg, err := c.ReadMessage(-1)
			if err == nil {
				fmt.Printf("Mensaje recibido: %s\n", string(msg.Value))
			} else {
				fmt.Printf("Error al leer el mensaje: %v\n", err)
			}
		}
	}

	// Cerrar consumidor de Kafka
	c.Close()
}
