package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gojou/goofmongo/pkg/connect"
	"github.com/gojou/goofmongo/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	p := models.Person{
		ID:        [12]byte{},
		Lastname:  "Poling",
		Firstname: "Rhi",
		Age:       13,
	}

	client := connect.GetClient()

	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	insertPerson(*client, p)
}
func insertPerson(c mongo.Client, p models.Person) {
	collection := c.Database("test").Collection("persons")
	insertResult, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted person with id: %v\n", insertResult.InsertedID)
}
