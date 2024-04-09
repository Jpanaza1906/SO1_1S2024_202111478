package main

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "server_grpc/proto"

	"google.golang.org/grpc"
)

// Context for the server
//var ctx = context.Background()

// Struct Server
type server struct {
	pb.UnimplementedGetInfoServer
}

// Const port
const (
	port = ":3001"
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
