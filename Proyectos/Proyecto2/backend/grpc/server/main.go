package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	pb "server_grpc/proto"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/grpc"
)

// Context for the server
//var ctx = context.Background()

// Struct Server
type server struct {
	pb.UnimplementedGetInfoServer
}

// Const port
var (
	port              = ":3001"
	kafkaBootstrapURL = "localhost:9092"
	kafkaTopic        = "topic-sopes1"
)

// Struct Data
type Data struct {
	Name  string `json:"name"`
	Album string `json:"album"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}

// Function to return data
func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	fmt.Println(">>> Received from client: ", in.GetName())
	data := Data{
		Name:  in.GetName(),
		Album: in.GetAlbum(),
		Year:  in.GetYear(),
		Rank:  in.GetRank(),
	}
	fmt.Println(">>> Data: ", data)

	// Aqui se mandaria la data a kafka

	// Convertir struct a JSON
	dataJSON, err := json.Marshal(data)
	if err != nil {
		log.Printf("Failed to marshal data to JSON: %v\n", err)
		return nil, err
	}

	// Configurar el productor de Kafka
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": kafkaBootstrapURL})
	if err != nil {
		log.Printf("Failed to create Kafka producer: %v\n", err)
		return nil, err
	}
	defer p.Close()

	// Enviar el mensaje a Kafka
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic, Partition: kafka.PartitionAny},
		Value:          []byte(dataJSON),
	}, nil)
	if err != nil {
		log.Printf("Failed to produce message to Kafka: %v\n", err)
		return nil, err
	}

	log.Println("Message sent to Kafka")

	return &pb.ReplyInfo{Info: "Hi Client! I received your data!"}, nil
}

// Main function
func main() {
	// Listen the port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf(">>> failed to listen: %v", err)
	}
	// Create a new server
	s := grpc.NewServer()
	// Register the server
	pb.RegisterGetInfoServer(s, &server{})
	// Serve the server
	if err := s.Serve(lis); err != nil {
		log.Fatalf(">>> failed to serve: %v", err)
	}
}
