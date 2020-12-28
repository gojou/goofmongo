package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type person struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`

	Lastname  string `bson:"lastname,omitempty" json:"lastname,omitempty"`
	Firstname string `bson:"firstname,omitempty" json:"firstname,omitempty"`
	Age       int    `bson:"age,omitempty" json:"age"`
}

func main() {

	client := getClient()

	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	p := person{
		//	ID:        [12]byte{},
		Lastname:  "Poling",
		Firstname: "Mark",
		Age:       57,
	}

	insertPerson(*client, p)
}
func insertPerson(c mongo.Client, p person) {
	collection := c.Database("test").Collection("persons")
	insertResult, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted person with id: %v\n", insertResult.InsertedID)
}

func getClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGOTESTURI"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
