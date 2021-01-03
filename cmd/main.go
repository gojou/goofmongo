package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gojou/goofmongo/pkg/models"
	"github.com/gojou/goofmongo/pkg/repo"
	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {

	p := models.Person{
		ID:        [12]byte{},
		Lastname:  "Tsung",
		Firstname: "Fred",
		Age:       55,
	}

	client := repo.GetClient()

	err := client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	insertPerson(*client, p)

	e := run()
	if e != nil {
		log.Fatalf("Fatal error: %v\n", e)
	}

}

func run() (e error) {
	r := mux.NewRouter()
	routes(r)

	// Critical to work on AppEngine
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
	return e
}

func insertPerson(c mongo.Client, p models.Person) {
	collection := c.Database("test").Collection("persons")
	insertResult, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted person with id: %v\n", insertResult.InsertedID)
}
