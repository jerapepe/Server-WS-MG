package mongo

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

type Message struct {
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	//create mongo client
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalf("Error creating mongo client: %v", err)
	}
	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatalf("Error pinging mongo: %v", err)
	}
	log.Println("Connected to MongoDB")
}

func GetMongoClient() *mongo.Client {
	return mongoClient
}

func SaveMessageToMongo(msg Message) bool {
	ctx := context.Background()
	collection := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))
	message := Message{
		Username:  msg.Username,
		Text:      msg.Text,
		Timestamp: time.Now(),
	}
	_, err := collection.InsertOne(ctx, message)
	if err != nil {
		log.Println("Error saving message to MongoDB Atlas: ", err)
		log.Fatal(err)
		return false
	}
	fmt.Println("Message saved to MongoDB Atlas")
	return true
}
