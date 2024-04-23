package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	pb "server_grpc/proto"

	"github.com/segmentio/kafka-go"
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
	port          = "localhost:3001"
	brokerAddress = "my-cluster-kafka-bootstrap:9092"
	kafkaTopic    = "topic-sopes1"
)

func produce(ctx context.Context, req *pb.RequestId) {
	l := log.New(os.Stdout, "Kafka Escribiendo: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   kafkaTopic,
		Logger:  l,
	})

	// Se crea el mensaje a enviar
	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte(req.Rank),
		Value: []byte(
			"{\"name\":\"" + req.Name +
				"\",\"album\":\"" + req.Album +
				"\",\"year\":\"" + req.Year +
				"\",\"rank\":\"" + req.Rank + "\"}",
		),
	})

	if err != nil {
		log.Fatal(">>> failed to write messages:", err)
	}
}

// Function to return data
func (s *server) ReturnInfo(ctx context.Context, req *pb.RequestId) (*pb.ReplyInfo, error) {
	// Se imprime la data recibida
	fmt.Println(">>> Received from client: ", req.Name, req.Album, req.Year, req.Rank)

	// Aqui se mandaria la data a kafka
	produce(ctx, req)

	return &pb.ReplyInfo{Info: "Hi Client! I received your data!"}, nil
}

// Main function
func main() {
	// Listen thmy-cluster-kafka-bootstrape port
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
