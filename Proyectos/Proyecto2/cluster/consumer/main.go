package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	kafkaTopic    = "topic-sopes1"
	brokerAddress = "my-cluster-kafka-bootstrap:9092"
)

var goRedisClient = redis.NewClient(&redis.Options{
	Addr: "redis-svc:6379",
})

type Data struct {
	Name  string `json:"name"`
	Album string `json:"album"`
	Year  string `json:"year"`
	Rank  string `json:"rank"`
}
type DataCont struct {
	Name     string `json:"name"`
	Album    string `json:"album"`
	Year     string `json:"year"`
	Rank     string `json:"rank"`
	Cantidad string `json:"cantidad"`
}
type DataDate struct {
	Name  string    `json:"name"`
	Album string    `json:"album"`
	Year  string    `json:"year"`
	Rank  string    `json:"rank"`
	Date  time.Time `json:"date"`
}

func consume(ctx context.Context) {
	l := log.New(os.Stdout, "Kafka Leyendo: ", 0)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   kafkaTopic,
		GroupID: "consumer-group",
		Logger:  l,
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatal(">>> failed to read message: ", err)
		}

		fmt.Println("Mensaje recibido: ", string(msg.Value))
		go insertRedis(string(msg.Value))
		go insertMongo(ctx, string(msg.Value))
	}
}

func insertRedis(jsonData string) error {
	var data Data
	json.Unmarshal([]byte(jsonData), &data)

	// Convertir mi objeto Data a DataCont
	dataCont := DataCont{
		Name:     data.Name,
		Album:    data.Album,
		Year:     data.Year,
		Rank:     data.Rank,
		Cantidad: "1",
	}

	// Consultar en redis si ya existe el objeto con key data.Rank

	ctx := context.Background()

	val, err := goRedisClient.Get(ctx, data.Rank).Result()

	// Si no existe, insertar el objeto con key data.Rank
	if err == redis.Nil {
		jsonData, _ := json.Marshal(dataCont)
		err := goRedisClient.Set(ctx, data.Rank, jsonData, 0).Err()
		if err != nil {
			log.Fatal(">>> failed to insert document: ", err)
		}
	} else {
		// Si ya existe, incrementar la cantidad
		var dataCont DataCont
		json.Unmarshal([]byte(val), &dataCont)

		// Convertir el valor a int , sumarle 1 y convertirlo a string
		cantidad, _ := strconv.Atoi(dataCont.Cantidad)
		cantidad++
		dataCont.Cantidad = strconv.Itoa(cantidad)

		jsonData, _ := json.Marshal(dataCont)
		err := goRedisClient.Set(ctx, data.Rank, jsonData, 0).Err()
		if err != nil {
			log.Fatal(">>> failed to insert document: ", err)
		}
	}

	return nil
}

func insertMongo(ctx context.Context, body string) {
	var data Data
	json.Unmarshal([]byte(body), &data)

	// Objeto DataDate
	dataDate := DataDate{
		Name:  data.Name,
		Album: data.Album,
		Year:  data.Year,
		Rank:  data.Rank,
		Date:  time.Now(),
	}

	var collection = GetCollection("bands")
	var err error
	_, err = collection.InsertOne(ctx, dataDate)

	if err != nil {
		log.Fatal(">>> failed to insert document: ", err)
	}
}

func GetCollection(collectionName string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://mongo-svc:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(">>> failed to connect to mongo: ", err)
	}

	collection := client.Database("so1p2").Collection(collectionName)
	return collection
}

func main() {
	ctx := context.Background()
	consume(ctx)
}
