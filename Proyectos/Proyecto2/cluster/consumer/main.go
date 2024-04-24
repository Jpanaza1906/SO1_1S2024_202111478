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

	// Incrementar el contador total
	ctx := context.Background()
	err := goRedisClient.Incr(ctx, "total").Err()
	if err != nil {
		log.Fatal(">>> failed to increment total counter: ", err)
	}

	// Crear la clave según el formato especificado
	key := fmt.Sprintf("%s (%s) - (%s) : No:%s", data.Album, data.Year, data.Name, data.Rank)

	// Consultar en Redis si ya existe un contador para la clave
	val, err := goRedisClient.HGet(ctx, "bands", key).Result()

	if err == redis.Nil {
		// Si no existe, insertar un nuevo contador para la clave
		err := goRedisClient.HSet(ctx, "bands", key, 1).Err()
		if err != nil {
			log.Fatal(">>> failed to insert document: ", err)
		}
	} else {
		// Si ya existe, incrementar el contador
		count, _ := strconv.Atoi(val)
		count++
		err := goRedisClient.HSet(ctx, "bands", key, count).Err()
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
	uri := "mongodb://jpanaza:so1p2@mongo-svc:27017/so1p2"
	clientOptions := options.Client().ApplyURI(uri)
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
