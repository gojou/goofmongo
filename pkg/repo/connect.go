package repo

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetClient gets a mongo client
func GetClient() *mongo.Client {
	fmt.Println(os.Getenv("MONGOTESTURI"))

	clientOptions := options.Client().ApplyURI("mongodb+srv://mongouser:hpkns372@cluster0.dfch4.mongodb.net/test")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Bad URI: ", err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
