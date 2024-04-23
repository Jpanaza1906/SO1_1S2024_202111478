package main

import (
	pb "client_grpc/proto"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Context for the client
var ctx = context.Background()

// Struct Data
type Data struct {
	Name  string `json:"name"`
	Album string `json:"album"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}

// Function to insert data
func insertData(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	voto := Data{
		Name:  data["name"],
		Album: data["album"],
		Year:  data["year"],
		Rank:  data["rank"],
	}

	go sendServer(voto)
	return nil
}

// Function to send data to the server
func sendServer(voto Data) {
	conn, err := grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf(">>> did not connect: %v", err)
	}

	cl := pb.NewGetInfoClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf(">>> did not close: %v", err)
		}
	}(conn)

	ret, err := cl.ReturnInfo(ctx, &pb.RequestId{
		Name:  voto.Name,
		Album: voto.Album,
		Year:  voto.Year,
		Rank:  voto.Rank,
	})

	if err != nil {
		log.Fatalf(">>> could not greet: %v", err)
	}

	fmt.Println(">>> Response from server: ", ret)

}

// API main function
func main() {
	app := fiber.New()

	app.Get("/grpc", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"res": "API is working",
		})
	})

	app.Post("/grpc/insert", insertData)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf(">>> did not listen: %v", err)
	}
}
